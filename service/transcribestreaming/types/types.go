// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
)

// A list of possible alternative transcriptions for the input audio. Each
// alternative may contain one or more of Items, Entities, or Transcript.
type Alternative struct {

	// Contains entities identified as personally identifiable information (PII) in
	// your transcription output.
	Entities []Entity

	// Contains words, phrases, or punctuation marks in your transcription output.
	Items []Item

	// Contains transcribed text.
	Transcript *string

	noSmithyDocumentSerde
}

// A wrapper for your audio chunks. Your audio stream consists of one or more audio
// events, which consist of one or more audio chunks. For more information, see
// Event stream encoding
// (https://docs.aws.amazon.com/transcribe/latest/dg/event-stream.html).
type AudioEvent struct {

	// An audio blob that contains the next part of the audio that you want to
	// transcribe. The maximum audio chunk size is 32 KB.
	AudioChunk []byte

	noSmithyDocumentSerde
}

// An encoded stream of audio blobs. Audio streams are encoded as either HTTP/2 or
// WebSocket data frames. For more information, see Transcribing streaming audio
// (https://docs.aws.amazon.com/transcribe/latest/dg/streaming.html).
//
// The following types satisfy this interface:
//
//	AudioStreamMemberAudioEvent
type AudioStream interface {
	isAudioStream()
}

// A blob of audio from your application. Your audio stream consists of one or more
// audio events. For more information, see Event stream encoding
// (https://docs.aws.amazon.com/transcribe/latest/dg/event-stream.html).
type AudioStreamMemberAudioEvent struct {
	Value AudioEvent

	noSmithyDocumentSerde
}

func (*AudioStreamMemberAudioEvent) isAudioStream() {}

// Contains entities identified as personally identifiable information (PII) in
// your transcription output, along with various associated attributes. Examples
// include category, confidence score, type, stability score, and start and end
// times.
type Entity struct {

	// The category of information identified. The only category is PII.
	Category *string

	// The confidence score associated with the identified PII entity in your audio.
	// Confidence scores are values between 0 and 1. A larger value indicates a higher
	// probability that the identified entity correctly matches the entity spoken in
	// your media.
	Confidence *float64

	// The word or words identified as PII.
	Content *string

	// The end time, in milliseconds, of the utterance that was identified as PII.
	EndTime float64

	// The start time, in milliseconds, of the utterance that was identified as PII.
	StartTime float64

	// The type of PII identified. For example, NAME or CREDIT_DEBIT_NUMBER.
	Type *string

	noSmithyDocumentSerde
}

// A word, phrase, or punctuation mark in your transcription output, along with
// various associated attributes, such as confidence score, type, and start and end
// times.
type Item struct {

	// The confidence score associated with a word or phrase in your transcript.
	// Confidence scores are values between 0 and 1. A larger value indicates a higher
	// probability that the identified item correctly matches the item spoken in your
	// media.
	Confidence *float64

	// The word or punctuation that was transcribed.
	Content *string

	// The end time, in milliseconds, of the transcribed item.
	EndTime float64

	// If speaker partitioning is enabled, Speaker labels the speaker of the specified
	// item.
	Speaker *string

	// If partial result stabilization is enabled, Stable indicates whether the
	// specified item is stable (true) or if it may change when the segment is complete
	// (false).
	Stable *bool

	// The start time, in milliseconds, of the transcribed item.
	StartTime float64

	// The type of item identified. Options are: PRONUNCIATION (spoken words) and
	// PUNCTUATION.
	Type ItemType

	// Indicates whether the specified item matches a word in the vocabulary filter
	// included in your request. If true, there is a vocabulary filter match.
	VocabularyFilterMatch bool

	noSmithyDocumentSerde
}

// The language code that represents the language identified in your audio,
// including the associated confidence score. If you enabled channel identification
// in your request and each channel contained a different language, you will have
// more than one LanguageWithScore result.
type LanguageWithScore struct {

	// The language code of the identified language.
	LanguageCode LanguageCode

	// The confidence score associated with the identified language code. Confidence
	// scores are values between zero and one; larger values indicate a higher
	// confidence in the identified language.
	Score float64

	noSmithyDocumentSerde
}

// A list of possible alternative transcriptions for the input audio. Each
// alternative may contain one or more of Items, Entities, or Transcript.
type MedicalAlternative struct {

	// Contains entities identified as personal health information (PHI) in your
	// transcription output.
	Entities []MedicalEntity

	// Contains words, phrases, or punctuation marks in your transcription output.
	Items []MedicalItem

	// Contains transcribed text.
	Transcript *string

	noSmithyDocumentSerde
}

// Contains entities identified as personal health information (PHI) in your
// transcription output, along with various associated attributes. Examples include
// category, confidence score, type, stability score, and start and end times.
type MedicalEntity struct {

	// The category of information identified. The only category is PHI.
	Category *string

	// The confidence score associated with the identified PHI entity in your audio.
	// Confidence scores are values between 0 and 1. A larger value indicates a higher
	// probability that the identified entity correctly matches the entity spoken in
	// your media.
	Confidence *float64

	// The word or words identified as PHI.
	Content *string

	// The end time, in milliseconds, of the utterance that was identified as PHI.
	EndTime float64

	// The start time, in milliseconds, of the utterance that was identified as PHI.
	StartTime float64

	noSmithyDocumentSerde
}

// A word, phrase, or punctuation mark in your transcription output, along with
// various associated attributes, such as confidence score, type, and start and end
// times.
type MedicalItem struct {

	// The confidence score associated with a word or phrase in your transcript.
	// Confidence scores are values between 0 and 1. A larger value indicates a higher
	// probability that the identified item correctly matches the item spoken in your
	// media.
	Confidence *float64

	// The word or punctuation that was transcribed.
	Content *string

	// The end time, in milliseconds, of the transcribed item.
	EndTime float64

	// If speaker partitioning is enabled, Speaker labels the speaker of the specified
	// item.
	Speaker *string

	// The start time, in milliseconds, of the transcribed item.
	StartTime float64

	// The type of item identified. Options are: PRONUNCIATION (spoken words) and
	// PUNCTUATION.
	Type ItemType

	noSmithyDocumentSerde
}

// The Result associated with a . Contains a set of transcription results from one
// or more audio segments, along with additional information per your request
// parameters. This can include information relating to alternative transcriptions,
// channel identification, partial result stabilization, language identification,
// and other transcription-related data.
type MedicalResult struct {

	// A list of possible alternative transcriptions for the input audio. Each
	// alternative may contain one or more of Items, Entities, or Transcript.
	Alternatives []MedicalAlternative

	// Indicates the channel identified for the Result.
	ChannelId *string

	// The end time, in milliseconds, of the Result.
	EndTime float64

	// Indicates if the segment is complete. If IsPartial is true, the segment is not
	// complete. If IsPartial is false, the segment is complete.
	IsPartial bool

	// Provides a unique identifier for the Result.
	ResultId *string

	// The start time, in milliseconds, of the Result.
	StartTime float64

	noSmithyDocumentSerde
}

// The MedicalTranscript associated with a . MedicalTranscript contains Results,
// which contains a set of transcription results from one or more audio segments,
// along with additional information per your request parameters.
type MedicalTranscript struct {

	// Contains a set of transcription results from one or more audio segments, along
	// with additional information per your request parameters. This can include
	// information relating to alternative transcriptions, channel identification,
	// partial result stabilization, language identification, and other
	// transcription-related data.
	Results []MedicalResult

	noSmithyDocumentSerde
}

// The MedicalTranscriptEvent associated with a MedicalTranscriptResultStream.
// Contains a set of transcription results from one or more audio segments, along
// with additional information per your request parameters.
type MedicalTranscriptEvent struct {

	// Contains Results, which contains a set of transcription results from one or more
	// audio segments, along with additional information per your request parameters.
	// This can include information relating to alternative transcriptions, channel
	// identification, partial result stabilization, language identification, and other
	// transcription-related data.
	Transcript *MedicalTranscript

	noSmithyDocumentSerde
}

// Contains detailed information about your streaming session.
//
// The following types satisfy this interface:
//
//	MedicalTranscriptResultStreamMemberTranscriptEvent
type MedicalTranscriptResultStream interface {
	isMedicalTranscriptResultStream()
}

// The MedicalTranscriptEvent associated with a MedicalTranscriptResultStream.
// Contains a set of transcription results from one or more audio segments, along
// with additional information per your request parameters. This can include
// information relating to alternative transcriptions, channel identification,
// partial result stabilization, language identification, and other
// transcription-related data.
type MedicalTranscriptResultStreamMemberTranscriptEvent struct {
	Value MedicalTranscriptEvent

	noSmithyDocumentSerde
}

func (*MedicalTranscriptResultStreamMemberTranscriptEvent) isMedicalTranscriptResultStream() {}

// The Result associated with a . Contains a set of transcription results from one
// or more audio segments, along with additional information per your request
// parameters. This can include information relating to alternative transcriptions,
// channel identification, partial result stabilization, language identification,
// and other transcription-related data.
type Result struct {

	// A list of possible alternative transcriptions for the input audio. Each
	// alternative may contain one or more of Items, Entities, or Transcript.
	Alternatives []Alternative

	// Indicates the channel identified for the Result.
	ChannelId *string

	// The end time, in milliseconds, of the Result.
	EndTime float64

	// Indicates if the segment is complete. If IsPartial is true, the segment is not
	// complete. If IsPartial is false, the segment is complete.
	IsPartial bool

	// The language code that represents the language spoken in your audio stream.
	LanguageCode LanguageCode

	// The language code of the dominant language identified in your stream. If you
	// enabled channel identification and each channel of your audio contains a
	// different language, you may have more than one result.
	LanguageIdentification []LanguageWithScore

	// Provides a unique identifier for the Result.
	ResultId *string

	// The start time, in milliseconds, of the Result.
	StartTime float64

	noSmithyDocumentSerde
}

// The Transcript associated with a . Transcript contains Results, which contains a
// set of transcription results from one or more audio segments, along with
// additional information per your request parameters.
type Transcript struct {

	// Contains a set of transcription results from one or more audio segments, along
	// with additional information per your request parameters. This can include
	// information relating to alternative transcriptions, channel identification,
	// partial result stabilization, language identification, and other
	// transcription-related data.
	Results []Result

	noSmithyDocumentSerde
}

// The TranscriptEvent associated with a TranscriptResultStream. Contains a set of
// transcription results from one or more audio segments, along with additional
// information per your request parameters.
type TranscriptEvent struct {

	// Contains Results, which contains a set of transcription results from one or more
	// audio segments, along with additional information per your request parameters.
	// This can include information relating to alternative transcriptions, channel
	// identification, partial result stabilization, language identification, and other
	// transcription-related data.
	Transcript *Transcript

	noSmithyDocumentSerde
}

// Contains detailed information about your streaming session.
//
// The following types satisfy this interface:
//
//	TranscriptResultStreamMemberTranscriptEvent
type TranscriptResultStream interface {
	isTranscriptResultStream()
}

// Contains Transcript, which contains Results. The  object contains a set of
// transcription results from one or more audio segments, along with additional
// information per your request parameters.
type TranscriptResultStreamMemberTranscriptEvent struct {
	Value TranscriptEvent

	noSmithyDocumentSerde
}

func (*TranscriptResultStreamMemberTranscriptEvent) isTranscriptResultStream() {}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isAudioStream()                   {}
func (*UnknownUnionMember) isMedicalTranscriptResultStream() {}
func (*UnknownUnionMember) isTranscriptResultStream()        {}
