package models

type SearchResult struct {
	Type     string      
	Data     interface{} 
	Matching []string    
}