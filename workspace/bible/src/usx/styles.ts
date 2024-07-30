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

export type BookHeaderStyle = 'ide' // File encoding information
  | 'h' // Running header text for a book
  | 'h1' // Running header text (DEPRECATED)
  | 'h2' // Running header text, left side of page (DEPRECATED)
  | 'h3' // Running header text, right side of page (DEPRECATED)
  | 'toc1' // Long table of contents text
  | 'toc2' // Short table of contents text
  | 'toc3' // Book Abbreviation
  | 'toca1' // Alternative language long table of contents text
  | 'toca2' // Alternative language short table of contents text
  | 'toca3' // Alternative language book Abbreviation
  | 'rem' // Remark
  | 'usfm' // USFM markup version information (may occur if USX was generated from USFM)

export const BookHeaderStyles: BookHeaderStyle[] = ['ide', 'h', 'h1', 'h2', 'h3', 'toc1', 'toc2', 'toc3', 'toca1', 'toca2', 'toca3', 'rem', 'usfm']

export type BookTitleStyle = 'mt' // The main title of the book (if single level)
  | 'mt1' // The main title of the book (if multiple levels)
  | 'mt2' // A secondary title usually occurring before the main title
  | 'mt3' // A tertiary title occurring after the main title
  | 'mt4'
  | 'imt' // Introduction major title, level 1 (if single level)
  | 'imt1' // Introduction major title, level 1 (if multiple levels)
  | 'imt2' // Introduction major title, level 2
  | 'rem' // Remark

export const BookTitleStyles: BookTitleStyle[] = ['mt', 'mt1', 'mt2', 'mt3', 'mt4', 'imt', 'imt1', 'imt2', 'rem']

export type BookIntroductionStyle = 'imt' // Introduction major title, level 1 - (if single level)
  | 'imt1' // Introduction major title, level 1 (if multiple levels)
  | 'imt2' // Introduction major title, level 2
  | 'imt3' // Introduction major title, level 3
  | 'imt4' // Introduction major title, level 4 (usually within parenthesis)
  | 'ib' // Introduction blank line
  | 'ie' // Introduction ending marker
  | 'ili' // A list entry, level 1 (if single level)
  | 'ili1' // A list entry, level 1 (if multiple levels)
  | 'ili2' // A list entry, level 2
  | 'im' // Introduction prose paragraph, with no first line indent (may occur after poetry)
  | 'imi' // Introduction prose paragraph text, indented, with no first line indent
  | 'imq' // Introduction prose paragraph, quote from the body text, with no first line indent
  | 'io' // Introduction outline text, level 1 (if single level)
  | 'io1' // Introduction outline text, level 1 (if multiple levels)
  | 'io2' // Introduction outline text, level 2
  | 'io3' // Introduction outline text, level 3
  | 'io4' // Introduction outline text, level 4
  | 'iot' // Introduction outline title
  | 'ip' // Introduction prose paragraph
  | 'ipi' // Introduction prose paragraph, indented, with first line indent
  | 'ipq' // Introduction prose paragraph, quote from the body text
  | 'ipr' // Introduction prose paragraph, right aligned
  | 'iq' // Introduction poetry text, level 1 (if single level)
  | 'iq1' // Introduction poetry text, level 1 (if multiple levels)
  | 'iq2' // Introduction poetry text, level 2
  | 'iq3' // Introduction poetry text, level 3
  | 'is' // Introduction section heading, level 1 (if single level)
  | 'is1' // Introduction section heading, level 1 (if multiple levels)
  | 'is2' // Introduction section heading, level 2
  | 'imte' // Introduction major title at introduction end, level 1 (if single level)
  | 'imte1' // Introduction major title at introduction end, level 1 (if multiple levels)
  | 'imte2' // Introduction major title at introduction end, level 2
  | 'iex' // Introduction explanatory or bridge text (e.g. explanation of missing book in Short Old Testament)
  | 'rem' // Remark

export const BookIntroductionStyles: BookIntroductionStyle[] = ['imt', 'imt1', 'imt2', 'imt3', 'imt4', 'ib', 'ie', 'ili', 'ili1', 'ili2', 'im', 'imi', 'imq', 'io', 'io1', 'io2', 'io3', 'io4', 'iot', 'ip', 'ipi', 'ipq', 'ipr', 'iq', 'iq1', 'iq2', 'iq3', 'is', 'is1', 'is2', 'imte', 'imte1', 'imte2', 'iex', 'rem']

export type BookIntroductionEndTitleStyle = 'mt' // The main title of the book (if single level)
  | 'mt1' // The main title of the book (if multiple levels)
  | 'mt2' // A secondary title usually occurring before the main title
  | 'mt3' // A tertiary title occurring after the main title
  | 'mt4' // A small secondary title sometimes occuring within parentheses
  | 'imt' // Introduction major title, level 1 (if single level)
  | 'imt1' // Introduction major title, level 1 (if multiple levels)
  | 'imt2' // Introduction major title, level 2

export const BookIntroductionEndTitleStyles: BookIntroductionEndTitleStyle[] = ['mt', 'mt1', 'mt2', 'mt3', 'mt4', 'imt', 'imt1', 'imt2']

export type BookChapterLabelStyle = 'cl'

export const BookChapterLabelStyles: BookChapterLabelStyle[] = ['cl']

export type CrossReferenceStyle = 'x' | 'ex'

export const CrossReferenceStyles: CrossReferenceStyle[] = ['x', 'ex']

export type CrossReferenceCharStyle = 'xo' // The cross reference origin reference
  | 'xop' // Published cross reference origin text (origin reference that should appear in the published text)
  | 'xt' // The cross reference target reference(s), protocanon only
  | 'xta' // Cross reference target references added text
  | 'xk' // A cross reference keyword
  | 'xq' // A cross-reference quotation from the scripture text
  | 'xot' // Cross-reference target reference(s), Old Testament only
  | 'xnt' // Cross-reference target reference(s), New Testament only
  | 'xdc' // Cross-reference target reference(s), Deuterocanon only (DEPRECATED - use char@style dc)

export const CrossReferenceCharStyles: CrossReferenceCharStyle[] = ['xo', 'xop', 'xt', 'xta', 'xk', 'xq', 'xot', 'xnt', 'xdc']

export type IntroCharStyle = 'ior' // Introduction references range for outline entry; for marking references separately
  | 'iqt' // For quoted scripture text appearing in the introduction

export const IntroCharStyles: IntroCharStyle[] = ['ior', 'iqt']

export type CharStyle = 'va' // Second (alternate) verse number (for coding dual numeration in Psalms; see also NRSV Exo 22.1-4)
  | 'vp' // Published verse marker - this is a verse marking that would be used in the published text
  | 'ca' // Second (alternate) chapter number
  | 'qac' // Poetry text, Acrostic markup of the first character of a line of acrostic poetry
  | 'qs' // Poetry text, Selah
  | 'add' // For a translational addition to the text
  | 'addpn' // For chinese words to be dot underline & underline (DEPRECATED - used nested char@style pn)
  | 'bk' // For the quoted name of a book
  | 'dc' // Deuterocanonical/LXX additions or insertions in the Protocanonical text
  | 'efm' // Reference to caller of previous footnote in a study Bible
  | 'fm' // Reference to caller of previous footnote
  | 'k' // For a keyword
  | 'nd' // For name of deity
  | 'ndx' // A subject index text item
  | 'ord' // For the text portion of an ordinal number
  | 'pn' // For a proper name
  | 'png' // For a geographic proper name
  | 'pro' // For indicating pronunciation in CJK texts (DEPRECATED - used char@style rb)
  | 'qt' // For Old Testament quoted text appearing in the New Testament
  | 'rq' // A cross-reference indicating the source text for the preceding quotation.
  | 'sig' // For the signature of the author of an Epistle
  | 'sls' // To represent where the original text is in a secondary language or from an alternate text source
  | 'tl' // For transliterated words
  | 'wg' // A Greek Wordlist text item
  | 'wh' // A Hebrew wordlist text item
  | 'wa' // An Aramaic wordlist text item
  | 'wj' // For marking the words of Jesus
  | 'xt' // A target reference(s)
  | 'jmp' // For associating linking attributes to a span of text
  | 'no' // A character style, use normal text
  | 'it' // A character style, use italic text
  | 'bd' // A character style, use bold text
  | 'bdit' // A character style, use bold + italic text
  | 'em' // A character style, use emphasized text style
  | 'sc' // A character style, for small capitalization text
  | 'sup' // A character style, for superscript text. Typically for use in critical edition footnotes.

export const CharStyles: CharStyle[] = ['va', 'vp', 'ca', 'qac', 'qs', 'add', 'addpn', 'bk', 'dc', 'efm', 'fm', 'k', 'nd', 'ndx', 'ord', 'pn', 'png', 'pro', 'qt', 'rq', 'sig', 'sls', 'tl', 'wg', 'wh', 'wa', 'wj', 'xt', 'jmp', 'no', 'it', 'bd', 'bdit', 'em', 'sc', 'sup']

export type FootnoteStyle = 'f' | 'fe' | 'ef'

export const FootnoteStyles: FootnoteStyle[] = ['f', 'fe', 'ef']

export type FootnoteCharStyle = 'fr' // The origin reference for the footnote
  | 'cat' // Note category (study Bible)
  | 'ft' // Footnote text, Protocanon
  | 'fk' // A footnote keyword
  | 'fq' // A footnote scripture quote or alternate rendering
  | 'fqa' // A footnote alternate rendering for a portion of scripture text
  | 'fl' // A footnote label text item, for marking or "labelling" the type or alternate translation being provided in the note.
  | 'fw' // A footnote witness list, for distinguishing a list of sigla representing witnesses in critical editions.
  | 'fp' // A Footnote additional paragraph marker
  | 'fv' // A verse number within the footnote text
  | 'fdc' // Footnote text, applies to Deuterocanon only (DEPRECATED - use char@style dc)
  | 'xt' // A cross reference target reference(s)
  | 'it' // A character style, use italic text
  | 'bd' // A character style, use bold text
  | 'bdit' // A character style, use bold + italic text
  | 'em' // A character style, use emphasized text style
  | 'sc' // A character style, for small capitalization text; abbreviations

export const FootnoteCharStyles: FootnoteCharStyle[] = ['fr', 'cat', 'ft', 'fk', 'fq', 'fqa', 'fl', 'fw', 'fp', 'fv', 'fdc', 'xt', 'it', 'bd', 'bdit', 'em', 'sc']

export type FootnoteVerseStyle = 'fv'

export const FootnoteVerseStyles: FootnoteVerseStyle[] = ['fv']

export type SidebarStyle = 'esb'

export const SidebarStyles: SidebarStyle[] = ['esb']

export type ListStyle = 'lh' // List header (introductory remark)
  | 'li' // A list entry, level 1 (if single level)
  | 'li1' // A list entry, level 1 (if multiple levels)
  | 'li2' // A list entry, level 2
  | 'li3' // A list entry, level 3
  | 'li4' // A list entry, level 4
  | 'lf' // List footer (introductory remark)
  | 'lim' // An embedded list entry, level 1 (if single level)
  | 'lim1' // An embedded list entry, level 1 (if multiple levels)
  | 'lim2' // An embedded list entry, level 2
  | 'lim3' // An embedded list entry, level 3
  | 'lim4' // An embedded list entry, level 4

export const ListStyles: ListStyle[] = ['lh', 'li', 'li1', 'li2', 'li3', 'li4', 'lf', 'lim', 'lim1', 'lim2', 'lim3', 'lim4']

export type ParaStyle =
  'restore' // Comment about when text was restored
  | 'cls' // Closure of an Epistle
  | 'iex' // Introduction explanatory or bridge text (e.g. explanation of missing book in Short Old Testament)
  | 'ip' // Division or Section introductory paragraph (study Bible)
  | 'lit' // For a comment or note inserted for liturgical use
  | 'm' //  Paragraph text, with no first line indent (may occur after poetry) aka: Paragraph Continuation
  | 'mi' // Paragraph text, indented, with no first line indent; often used for discourse
  | 'nb' // Paragraph text, with no break from previous paragraph text (at chapter boundary)
  | 'p' // Paragraph text, with first line indent
  | 'pb' // Page Break used for new reader portions and children's bibles where content is controlled by the page
  | 'pc' // Inscription (paragraph text centered)
  | 'pi' // Paragraph text, level 1 indent (if single level), with first line indent; often used for discourse
  | 'pi1' // Paragraph text, level 1 indent (if multiple levels), with first line indent; often used for discourse
  | 'pi2' // Paragraph text, level 2 indent, with first line indent; often used for discourse
  | 'pi3' // Paragraph text, level 3 indent, with first line indent; often used for discourse
  | 'po' // Letter opening
  | 'pr' // Text refrain (paragraph right-aligned)
  | 'pmo' // Embedded text opening
  | 'pm' // Embedded text paragraph
  | 'pmc' // Embedded text closing
  | 'pmr' // Embedded text refrain
  | 'ph' // Paragraph text, with level 1 hanging indent (if single level) (DEPRECATED - use para@style li//)
  | 'ph1' // Paragraph text, with level 1 hanging indent (if multiple levels)
  | 'ph2' // Paragraph text, with level 2 hanging indent
  | 'ph3' // Paragraph text, with level 3 hanging indent
  | 'q' // Poetry text, level 1 indent (if single level)
  | 'q1' // Poetry text, level 1 indent (if multiple levels)
  | 'q2' // Poetry text, level 2 indent
  | 'q3' // Poetry text, level 3 indent
  | 'q4' // Poetry text, level 4 indent
  | 'qa' // Poetry text, Acrostic marker/heading
  | 'qc' // Poetry text, centered
  | 'qr' // Poetry text, Right Aligned
  | 'qm' // Poetry text, embedded, level 1 indent (if single level)
  | 'qm1' // Poetry text, embedded, level 1 indent (if multiple levels)
  | 'qm2' // Poetry text, embedded, level 2 indent
  | 'qm3' // Poetry text, embedded, level 3 indent
  | 'qd' // A Hebrew musical performance annotation, similar in content to Hebrew descriptive title.
  | 'b' // Poetry text stanza break (e.g. stanza break)
  | 'd' // A Hebrew text heading, to provide description (e.g. Psalms)
  | 'ms' // A major section division heading, level 1 (if single level)
  | 'ms1' // A major section division heading, level 1 (if multiple levels)
  | 'ms2' // A major section division heading, level 2
  | 'ms3' // A major section division heading, level 3
  | 'mr' // A major section division references range heading
  | 'r' // Parallel reference(s)
  | 's' // A section heading, level 1 (if single level)
  | 's1' // A section heading, level 1 (if multiple levels)
  | 's2' // A section heading, level 2 (e.g. Proverbs 22-24)
  | 's3' // A section heading, level 3 (e.g. Genesis "The First Day")
  | 's4' // A section heading, level 4
  | 'sr' // A section division references range heading
  | 'sp' // A heading, to identify the speaker (e.g. Job)
  | 'sd' // Vertical space used to divide the text into sections, level 1 (if single level)
  | 'sd1' // Semantic division location (vertical space used to divide the text into sections), level 1 (if multiple levels)
  | 'sd2' // Semantic division location (vertical space used to divide the text into sections), level 2
  | 'sd3' // Semantic division location (vertical space used to divide the text into sections), level 3
  | 'sd4' // Semantic division location (vertical space used to divide the text into sections), level 4
  | 'ts' // Translator's chunk (to identify chunks of text suitable for translating at one time)
  | 'cp' // Published chapter number
  | 'cl' // Chapter label used for translations that add a word such as "Chapter"
  | 'cd' // Chapter Description (Publishing option D, e.g. in Russian Bibles)
  | 'mte' // The main title of the book repeated at the end of the book, level 1 (if single level)
  | 'mte1' // The main title of the book repeat /ed at the end of the book, level 1 (if multiple levels)
  | 'mte2' // A secondary title occurring before or after the 'ending' main title
  | 'p' // Front or back matter text paragraph, level 1
  | 'p1' // Front or back matter text paragraph, level 1 (if multiple levels)
  | 'p2' // Front or back matter text paragraph, level 2 (if multiple levels)
  | 'k1' // Concordance main entry text or keyword, level 1
  | 'k2' // Concordance main entry text or keyword, level 2
  | 'rem' // Remark

export const ParaStyles: ParaStyle[] = ['restore', 'cls', 'iex', 'ip', 'lit', 'm', 'mi', 'nb', 'p', 'pb', 'pc', 'pi', 'pi1', 'pi2', 'pi3', 'po', 'pr', 'pmo', 'pm', 'pmc', 'pmr', 'ph', 'ph1', 'ph2', 'ph3', 'q', 'q1', 'q2', 'q3', 'q4', 'qa', 'qc', 'qr', 'qm', 'qm1', 'qm2', 'qm3', 'qd', 'b', 'd', 'ms', 'ms1', 'ms2', 'ms3', 'mr', 'r', 's', 's1', 's2', 's3', 's4', 'sr', 'sp', 'sd', 'sd1', 'sd2', 'sd3', 'sd4', 'ts', 'cp', 'cl', 'cd', 'mte', 'mte1', 'mte2', 'p', 'p1', 'p2', 'k1', 'k2', 'rem']

export type VerseStartStyle = 'v'

export const VerseStartStyles: VerseStartStyle[] = ['v']

export type ListCharStyle = 'litl' // List entry total text
  | 'lik' // Structured list entry key text
  | 'liv' // Structured list entry value 1 content (if single value)
  | 'liv1' // Structured list entrt value 1 content (if multiple values)
  | 'liv2' // Structured list entry value 2 content
  | 'liv3' // Structured list entry value 3 content
  | 'liv4' // Structured list entry value 4 content
  | 'liv5' // Structured list entry value 5 content

export const ListCharStyles: ListCharStyle[] = ['litl', 'lik', 'liv', 'liv1', 'liv2', 'liv3', 'liv4', 'liv5']