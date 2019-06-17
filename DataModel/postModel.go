package DataModel

type Post struct {
	Id               int
	Asset            Asset
	Images           []string
	Comments         []Comment
	NegativeComments []NegativeComment
	Name             string
	Likes            []User
	Rank             int
	Date             string
}

func NewPost(id int, name string) *Post {
	return &Post{Id: id, Name: name}
}

//add new comment to post
func AddComment(post Post, comment Comment) {
	post.Comments = append(post.Comments, comment)
}

//add negative new comment to post
func AddNegativeComment(post Post, comment NegativeComment) {
	post.NegativeComments = append(post.NegativeComments, comment)
}

//add like to this post by specific user
func AddLike(post Post, user User) {
	post.Likes = append(post.Likes, user)
}

func SetRank(post Post) {
	post.Rank = post.Asset.User.Reliable_grade - len(post.NegativeComments)
}
