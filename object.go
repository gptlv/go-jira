package jira

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ObjectService struct {
	client *Client
}

type Object struct {
	ID           int          `json:"id"`
	Label        string       `json:"label"`
	ObjectKey    string       `json:"objectKey"`
	Avatar       Avatar       `json:"avatar"`
	ObjectType   ObjectType   `json:"objectType"`
	Created      string       `json:"created"`
	Updated      string       `json:"updated"`
	HasAvatar    bool         `json:"hasAvatar"`
	Timestamp    int64        `json:"timestamp"`
	Attributes   []Attribute  `json:"attributes"`
	ExtendedInfo ExtendedInfo `json:"extendedInfo"`
	Links        Links        `json:"_links"`
	Name         string       `json:"name"`
}
type Avatar struct {
	URL16    string `json:"url16"`
	URL48    string `json:"url48"`
	URL72    string `json:"url72"`
	URL144   string `json:"url144"`
	URL288   string `json:"url288"`
	ObjectID int    `json:"objectId"`
}
type Icon struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	URL16 string `json:"url16"`
	URL48 string `json:"url48"`
}
type ObjectType struct {
	ID                        int    `json:"id"`
	Name                      string `json:"name"`
	Type                      int    `json:"type"`
	Description               string `json:"description"`
	Icon                      Icon   `json:"icon"`
	Position                  int    `json:"position"`
	Created                   string `json:"created"`
	Updated                   string `json:"updated"`
	ObjectCount               int    `json:"objectCount"`
	ParentObjectTypeID        int    `json:"parentObjectTypeId"`
	ObjectSchemaID            int    `json:"objectSchemaId"`
	Inherited                 bool   `json:"inherited"`
	AbstractObjectType        bool   `json:"abstractObjectType"`
	ParentObjectTypeInherited bool   `json:"parentObjectTypeInherited"`
}
type DefaultType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type ObjectTypeAttribute struct {
	ID                      int         `json:"id"`
	Name                    string      `json:"name"`
	Label                   bool        `json:"label"`
	Type                    int         `json:"type"`
	DefaultType             DefaultType `json:"defaultType"`
	Editable                bool        `json:"editable"`
	System                  bool        `json:"system"`
	Sortable                bool        `json:"sortable"`
	Summable                bool        `json:"summable"`
	Indexed                 bool        `json:"indexed"`
	MinimumCardinality      int         `json:"minimumCardinality"`
	MaximumCardinality      int         `json:"maximumCardinality"`
	Removable               bool        `json:"removable"`
	Hidden                  bool        `json:"hidden"`
	IncludeChildObjectTypes bool        `json:"includeChildObjectTypes"`
	UniqueAttribute         bool        `json:"uniqueAttribute"`
	Options                 string      `json:"options"`
	Position                int         `json:"position"`
}
type ObjectAttributeValue struct {
	Value          string `json:"value"`
	DisplayValue   string `json:"displayValue"`
	SearchValue    string `json:"searchValue"`
	ReferencedType bool   `json:"referencedType"`
}

type ReferenceType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
	URL16       string `json:"url16"`
	Removable   bool   `json:"removable"`
}
type ReferenceObjectType struct {
	ID                        int    `json:"id"`
	Name                      string `json:"name"`
	Type                      int    `json:"type"`
	Description               string `json:"description"`
	Icon                      Icon   `json:"icon"`
	Position                  int    `json:"position"`
	Created                   string `json:"created"`
	Updated                   string `json:"updated"`
	ObjectCount               int    `json:"objectCount"`
	ParentObjectTypeID        int    `json:"parentObjectTypeId"`
	ObjectSchemaID            int    `json:"objectSchemaId"`
	Inherited                 bool   `json:"inherited"`
	AbstractObjectType        bool   `json:"abstractObjectType"`
	ParentObjectTypeInherited bool   `json:"parentObjectTypeInherited"`
}

type Attribute struct {
	ID                    int                    `json:"id"`
	ObjectTypeAttribute   ObjectTypeAttribute    `json:"objectTypeAttribute,omitempty"`
	ObjectTypeAttributeID int                    `json:"objectTypeAttributeId"`
	ObjectAttributeValues []ObjectAttributeValue `json:"objectAttributeValues"`
	ObjectID              int                    `json:"objectId"`
}
type ExtendedInfo struct {
	OpenIssuesExists  bool `json:"openIssuesExists"`
	AttachmentsExists bool `json:"attachmentsExists"`
}
type Links struct {
	Self string `json:"self"`
}

type GetObjectQueryOptions struct {
	IncludeAttributes   bool `url:"includeAttributes,omitempty"`
	IncludeExtendedInfo bool `url:"includeExtendedInfo,omitempty"`
}

type UpdateObjectPayload struct {
	ObjectTypeId string      `json:"objectTypeId,omitempty"`
	Attributes   []Attribute `json:"attributes,omitempty"`
}

type FindObjectPayload struct {
	Query                  string               `json:"qlQuery,omitempty"`
	ObjectTypeID           string               `json:"objectTypeId,omitempty"`
	Page                   int                  `json:"page,omitempty"`
	ResultPerPage          int                  `json:"resultsPerPage,omitempty"`
	OrderByTypeAttributeID int                  `json:"orderByTypeAttrId,omitempty"`
	Asc                    int                  `json:"asc,omitempty"`
	ObjectSchemaID         int                  `json:"objectSchemaId,omitempty"`
	IncludeAttributes      bool                 `json:"includeAttributes,omitempty"`
	AttributesToDisplay    *AttributesToDisplay `json:"attributesToDisplay,omitempty"`
}

type AttributesToDisplay struct {
	AttributesToDisplayIDs []int `json:"attributesToDisplayIds,omitempty"`
}

type ObjectList struct {
	ObjectEntries         []Object              `json:"objectEntries"`
	ObjectTypeAttributes  []ObjectTypeAttribute `json:"objectTypeAttributes"`
	ObjectTypeID          int                   `json:"objectTypeId"`
	ObjectTypeIsInherited bool                  `json:"objectTypeIsInherited"`
	AbstractObjectType    bool                  `json:"abstractObjectType"`
	TotalFilterCount      int                   `json:"totalFilterCount"`
	StartIndex            int                   `json:"startIndex"`
	ToIndex               int                   `json:"toIndex"`
	PageObjectSize        int                   `json:"pageObjectSize"`
	PageNumber            int                   `json:"pageNumber"`
	OrderWay              string                `json:"orderWay"`
	Filters               []Filters             `json:"filters"`
	QlQuery               string                `json:"qlQuery"`
	QlQuerySearchResult   bool                  `json:"qlQuerySearchResult"`
	ConversionPossible    bool                  `json:"conversionPossible"`
	MatchedFilterValues   []MatchedFilterValues `json:"matchedFilterValues"`
	// InheritanceTree       InheritanceTree       `json:"inheritanceTree"`
	Iql             string `json:"iql"`
	IqlSearchResult bool   `json:"iqlSearchResult"`
	PageSize        int    `json:"pageSize"`
}

type Filters struct {
	ObjectTypeAttributeID int      `json:"objectTypeAttributeId"`
	SelectedValues        []string `json:"selectedValues"`
}

type MatchedFilterValues struct {
	ID                    int                    `json:"id"`
	ObjectTypeAttribute   ObjectTypeAttribute    `json:"objectTypeAttribute"`
	ObjectTypeAttributeID int                    `json:"objectTypeAttributeId"`
	ObjectAttributeValues []ObjectAttributeValue `json:"objectAttributeValues"`
	ObjectID              int                    `json:"objectId"`
}

// https://docs.atlassian.com/assets/REST/9.1.16/#object-loadObject
func (s *ObjectService) GetWithContext(ctx context.Context, objectID string, options *GetObjectQueryOptions) (*Object, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/insight/1.0/object/%s", objectID)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodGet, apiEndPoint, nil)
	if err != nil {
		return nil, nil, err
	}

	if options != nil {
		q, err := query.Values(options)
		if err != nil {
			return nil, nil, err
		}
		req.URL.RawQuery = q.Encode()
	}

	object := new(Object)
	resp, err := s.client.Do(req, object)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return object, resp, nil
}

func (s *ObjectService) Get(objectID string, options *GetObjectQueryOptions) (*Object, *Response, error) {
	return s.GetWithContext(context.Background(), objectID, options)
}

// https://docs.atlassian.com/assets/REST/9.1.16/#object-updateObject
func (s *ObjectService) UpdateWithContext(ctx context.Context, objectID string, payload *UpdateObjectPayload) (*Object, *Response, error) {
	apiEndPoint := fmt.Sprintf("rest/insight/1.0/object/%s", objectID)
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPut, apiEndPoint, payload)
	if err != nil {
		return nil, nil, err
	}

	object := new(Object)
	resp, err := s.client.Do(req, object)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return object, resp, nil
}

func (s *ObjectService) Update(objectID string, payload *UpdateObjectPayload) (*Object, *Response, error) {
	return s.UpdateWithContext(context.Background(), objectID, payload)
}

// https://docs.atlassian.com/assets/REST/9.1.16/#object-findObject
func (s *ObjectService) FindWithContext(ctx context.Context, payload *FindObjectPayload) (*ObjectList, *Response, error) {
	apiEndPoint := "rest/insight/1.0/object/navlist/iql"
	req, err := s.client.NewRequestWithContext(ctx, http.MethodPost, apiEndPoint, payload)
	if err != nil {
		return nil, nil, err
	}

	objectList := new(ObjectList)
	resp, err := s.client.Do(req, objectList)
	if err != nil {
		jerr := NewJiraError(resp, err)
		return nil, resp, jerr
	}

	return objectList, resp, nil
}

func (s *ObjectService) Find(payload *FindObjectPayload) (*ObjectList, *Response, error) {
	return s.FindWithContext(context.Background(), payload)
}
