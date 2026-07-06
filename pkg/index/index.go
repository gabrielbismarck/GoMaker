package index

import (
    "bufio"
    "bytes"
    "encoding/gob"
    "errors"
    "fmt"
    "log"
    "os"
    "strings"
    "sync"
)

type InvertedIndex map[string]map[string]int

type DocumentFrequency map[string]int

type SavedIndex struct {
    Index     InvertedIndex
    DocFreq   DocumentFrequency
    Documents []string // Lista de URLs de documentos indexados
}

type Indexer struct {
    Index     InvertedIndex
    DocFreq   DocumentFrequency
    Documents []string
    mu        sync.RWMutex
    name      string       
}

func NewIndexer(name string) *Indexer {
    idx := &Indexer{
        name: name,
    }
    idx.createEmptyIndex()

    if err := idx.loadIndex(); err != nil {
        if errors.Is(err, os.ErrNotExist) {
            log.Printf("Arquivo de índice '%s.db' não encontrado. Criando um novo índice vazio.", name)
        } else {
            log.Printf("Erro ao carregar o índice de '%s.db': %v. Iniciando com um índice vazio.", name, err)
        }
    } else {
        log.Printf("Índice carregado com sucesso de '%s.db'.", name)
    }

    return idx
}

func (idx *Indexer) createEmptyIndex() {
    idx.Index = make(InvertedIndex)
    idx.DocFreq = make(DocumentFrequency)
    idx.Documents = make([]string, 0)
}

func (idx *Indexer) AddDocToIndex(url string, content string) {
    idx.mu.Lock() 
    defer idx.mu.Unlock()

    idx.Documents = append(idx.Documents, url)

    reader := strings.NewReader(content)
    seenTerms := make(map[string]bool)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        word := strings.ToLower(strings.Trim(scanner.Text(), ",.!?&<>;:=§$%{}[]()|"))

        if idx.Index[word] == nil {
            idx.Index[word] = make(map[string]int)
        }
        idx.Index[word][url]++

        if !seenTerms[word] {
            idx.DocFreq[word]++
            seenTerms[word] = true
        }
    }
}

func (idx *Indexer) SaveIndex() error {
    idx.mu.RLock()
    defer idx.mu.RUnlock()

    savedIndex := SavedIndex{
        Index:     idx.Index,
        DocFreq:   idx.DocFreq,
        Documents: idx.Documents,
    }

    indexFilePath := fmt.Sprintf("/tmp/%s.db", idx.name)
    err := idx.writeStructToFile(indexFilePath, savedIndex)
    if err != nil {
        return fmt.Errorf("erro ao salvar o índice: %w", err)
    }
    log.Printf("Índice salvo com sucesso em '%s'.", indexFilePath)
    return nil
}

func (idx *Indexer) loadIndex() error {
    indexFilePath := fmt.Sprintf("/tmp/%s.db", idx.name)
    if _, err := os.Stat(indexFilePath); errors.Is(err, os.ErrNotExist) {
        return os.ErrNotExist
    }

    var savedIndex SavedIndex
    err := idx.readStructFromFile(indexFilePath, &savedIndex)
    if err != nil {
        return fmt.Errorf("erro ao ler o índice do arquivo: %w", err)
    }

    idx.Index = savedIndex.Index
    idx.DocFreq = savedIndex.DocFreq
    idx.Documents = savedIndex.Documents
    return nil
}

func (idx *Indexer) writeStructToFile(filename string, data interface{}) error {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(data)
    if err != nil {
        return fmt.Errorf("erro ao serializar as structs: %w", err)
    }

    err = os.WriteFile(filename, buf.Bytes(), 0644)
    if err != nil {
        return fmt.Errorf("erro ao gravar o arquivo '%s': %w", filename, err)
    }
    return nil
}

func (idx *Indexer) readStructFromFile(filename string, data interface{}) error {
    content, err := os.ReadFile(filename)
    if err != nil {
        return fmt.Errorf("erro ao ler o arquivo '%s': %w", filename, err)
    }

    buf := bytes.NewBuffer(content)
    dec := gob.NewDecoder(buf)
    err = dec.Decode(data)
    if err != nil {
        return fmt.Errorf("erro ao desserializar as structs do arquivo '%s': %w", filename, err)
    }
    return nil
}
