package model

type Post struct {
	Title string `json:"title"`
	Post  string `json:"message"`
	ID    int64  `json:"id"`
}

type Newsfeed struct {
	Posts map[int64]Post
	Count int64
}

func New() *Newsfeed {
	return &Newsfeed{
		Posts: make(map[int64]Post),
		Count: 0,
	}
}

type NewsfeedStore interface {
	GetAll(startId int64, endId int64) []Post
	Add(item Post)
	NextId() int64
	Update(item Post) bool
	Delete(id int64) bool
	Read(id int64) Post
}

func (r *Newsfeed) NextId() int64 {
	return r.Count + 1
}

func (r *Newsfeed) Add(item Post) {
	r.Posts[item.ID] = item
	r.Count = r.Count + 1
}

func (r *Newsfeed) Update(item Post) bool {
	if _, ok := r.Posts[item.ID]; ok {
		r.Posts[item.ID] = item
		return true
	}
	return false
}

func (r *Newsfeed) Read(id int64) Post {
	if val, ok := r.Posts[id]; ok {
		return val
	}
	return Post{}
}

func (r *Newsfeed) Delete(id int64) bool {
	if _, ok := r.Posts[id]; ok {
		delete(r.Posts, id)
		return true
	}
	return false
}

func (r *Newsfeed) GetAll(startId int64, endId int64) []Post {
	values := []Post{}

	for id, value := range r.Posts {
		if id >= startId && id <= endId {
			values = append(values, value)
		}
	}
	return values
}
