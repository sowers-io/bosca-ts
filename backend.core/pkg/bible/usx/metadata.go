/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package usx

type MetadataIdentification struct {
	Name              string `xml:"name"`
	NameLocal         string `xml:"nameLocal"`
	Description       string `xml:"description"`
	Abbreviation      string `xml:"abbreviation"`
	AbbreviationLocal string `xml:"abbreviationLocal"`
	Scope             string `xml:"scope"`
	DateCompleted     string `xml:"dateCompleted"`
	BundleProducer    string `xml:"bundleProducer"`
	SystemId          []struct {
		Type     string `xml:"type,attr"`
		ID       string `xml:"id"`
		CsetId   string `xml:"csetId"`
		Name     string `xml:"name"`
		FullName string `xml:"fullName"`
	} `xml:"systemId"`
}

type MetadataType struct {
	Medium          string `xml:"medium"`
	IsConfidential  string `xml:"isConfidential"`
	HasCharacters   string `xml:"hasCharacters"`
	IsTranslation   string `xml:"isTranslation"`
	IsExpression    string `xml:"isExpression"`
	TranslationType string `xml:"translationType"`
	Audience        string `xml:"audience"`
	ProjectType     string `xml:"projectType"`
}

type MetadataLanguage struct {
	Iso             string `xml:"iso"`
	Name            string `xml:"name"`
	NameLocal       string `xml:"nameLocal"`
	Script          string `xml:"script"`
	ScriptCode      string `xml:"scriptCode"`
	ScriptDirection string `xml:"scriptDirection"`
	Ldml            string `xml:"ldml"`
	Rod             string `xml:"rod"`
	Numerals        string `xml:"numerals"`
}

type MetadataName struct {
	ID    string `xml:"id,attr"`
	Abbr  string `xml:"abbr"`
	Short string `xml:"short"`
	Long  string `xml:"long"`
}

type MetadataRightsHolder struct {
	Abbr      string `xml:"abbr"`
	URL       string `xml:"url"`
	NameLocal string `xml:"nameLocal"`
	Uid       string `xml:"uid"`
	Name      string `xml:"name"`
}

type MetadataRightsAdmin struct {
	URL  string `xml:"url"`
	Uid  string `xml:"uid"`
	Name string `xml:"name"`
}

type MetadataContributor struct {
	Content     string `xml:"content"`
	Publication string `xml:"publication"`
	Management  string `xml:"management"`
	Finance     string `xml:"finance"`
	Qa          string `xml:"qa"`
	Uid         string `xml:"uid"`
	Name        string `xml:"name"`
}

type MetadataAgencies struct {
	RightsHolder MetadataRightsHolder `xml:"rightsHolder"`
	RightsAdmin  MetadataRightsAdmin  `xml:"rightsAdmin"`
	Contributor  MetadataContributor  `xml:"contributor"`
}

type MetadataCountry struct {
	Iso  string `xml:"iso"`
	Name string `xml:"name"`
}

type MetadataCountries struct {
	Country MetadataCountry `xml:"country"`
}

type MetadataBook struct {
	Code string `xml:"code,attr"`
}

type MetadataResource struct {
	Checksum string `xml:"checksum,attr"`
	MimeType string `xml:"mimeType,attr"`
	Size     string `xml:"size,attr"`
	URI      string `xml:"uri,attr"`
}

type MetadataStructureContent struct {
	Name string `xml:"name,attr"`
	Src  string `xml:"src,attr"`
	Role string `xml:"role,attr"`
}

type MetadataPublication struct {
	ID                string `xml:"id,attr"`
	Default           string `xml:"default,attr"`
	Name              string `xml:"name"`
	NameLocal         string `xml:"nameLocal"`
	Description       string `xml:"description"`
	DescriptionLocal  string `xml:"descriptionLocal"`
	Abbreviation      string `xml:"abbreviation"`
	AbbreviationLocal string `xml:"abbreviationLocal"`
	CanonicalContent  struct {
		Book []MetadataBook `xml:"book"`
	} `xml:"canonicalContent"`
	Structure struct {
		Content []MetadataStructureContent `xml:"content"`
	} `xml:"structure"`
}

type Metadata struct {
	ID             string                 `xml:"id,attr"`
	Version        string                 `xml:"version,attr"`
	Revision       string                 `xml:"revision,attr"`
	Identification MetadataIdentification `xml:"identification"`
	Type           MetadataType           `xml:"type"`
	Relationships  string                 `xml:"relationships"`
	Agencies       MetadataAgencies       `xml:"agencies"`
	Language       MetadataLanguage       `xml:"language"`
	Countries      MetadataCountries      `xml:"countries"`
	Format         struct {
		UsxVersion       string `xml:"usxVersion"`
		VersedParagraphs string `xml:"versedParagraphs"`
	} `xml:"format"`
	Names struct {
		Name []MetadataName `xml:"name"`
	} `xml:"names"`
	Manifest struct {
		Resource []MetadataResource `xml:"resource"`
	} `xml:"manifest"`
	Source struct {
		CanonicalContent struct {
			Book []MetadataBook `xml:"book"`
		} `xml:"canonicalContent"`
		Structure struct {
			Content struct {
				Src  string `xml:"src,attr"`
				Role string `xml:"role,attr"`
			} `xml:"content"`
		} `xml:"structure"`
	} `xml:"source"`
	Publications struct {
		Publication MetadataPublication `xml:"publication"`
	} `xml:"publications"`
	Copyright struct {
		Text          string `xml:",chardata"`
		FullStatement struct {
			Text             string `xml:",chardata"`
			StatementContent struct {
				Text string `xml:",chardata"`
				Type string `xml:"type,attr"`
				P    struct {
					Text   string `xml:",chardata"`
					Strong string `xml:"strong"`
				} `xml:"p"`
			} `xml:"statementContent"`
		} `xml:"fullStatement"`
	} `xml:"copyright"`
	Promotion struct {
		PromoVersionInfo struct {
			Text        string `xml:",chardata"`
			ContentType string `xml:"contentType,attr"`
			P           []struct {
				Text string `xml:",chardata"`
				A    struct {
					Text string `xml:",chardata"`
					Href string `xml:"href,attr"`
				} `xml:"a"`
			} `xml:"p"`
		} `xml:"promoVersionInfo"`
	} `xml:"promotion"`
	ArchiveStatus struct {
		Text          string `xml:",chardata"`
		ArchivistName string `xml:"archivistName"`
		DateArchived  string `xml:"dateArchived"`
		DateUpdated   string `xml:"dateUpdated"`
		Comments      string `xml:"comments"`
	} `xml:"archiveStatus"`
}
