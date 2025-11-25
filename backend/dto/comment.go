package dto

type CommentDTO struct {
	IDSubscription int    `json:"id"`
	Comment        string `json:"comment"`
}

type CommentsDTO []CommentDTO

type ResponseCommentDTO struct {
	IDSubscription int    `json:"id"`
	Comment        string `json:"comment"`
	Username       string `json:"username"`
}

type ResponseCommentsDTO []ResponseCommentDTO
