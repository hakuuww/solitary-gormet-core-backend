package models

//If the frontend uses form to submitt requests, we can add struct tags inside the struct
//Since the frontend is going to send data in JSON, we use JSON tags

/*
Communication with the Frontend:

    When your backend responds to a client request with a Video struct,
	it will be automatically serialized into a JSON representation using the field tags you provided.
	This JSON response can be sent to the frontend, which might be a web browser,
	mobile app, or any other client capable of processing JSON data.

    The frontend can then parse this JSON response and access the data by using the same field names defined in your struct.
	This consistency in field naming between the backend and frontend is essential for seamless communication.

Unmarshaling:

    When you receive JSON data from the frontend (e.g., in an HTTP request),
	Go can unmarshal it into a Video struct by matching the JSON keys to the struct field names defined in the tags.
	This allows you to work with the data in your Go code more easily.
*/
type Video struct {
	Id          int    `uri:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
