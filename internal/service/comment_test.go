package service

import (
	"brewnique.fdunlap.com/internal/data"
	"fmt"
	"testing"
	"time"
)

type TestCommentProvider struct {
	comments     map[int64]*data.Comment
	commentVotes map[int64]map[int64]*data.CommentVote // commentId -> userId -> vote
	nextID       int64
}

func NewTestCommentProvider() *TestCommentProvider {
	return &TestCommentProvider{
		comments:     make(map[int64]*data.Comment),
		commentVotes: make(map[int64]map[int64]*data.CommentVote),
		nextID:       1,
	}
}

func (p *TestCommentProvider) PutComment(comment *data.Comment) (data.Comment, error) {
	comment.Id = p.nextID
	comment.CreatedAt = time.Now()
	comment.UpdatedAt = time.Now()
	p.commentVotes[comment.Id] = make(map[int64]*data.CommentVote)
	p.comments[comment.Id] = comment
	p.nextID++
	return *comment, nil
}

func (p *TestCommentProvider) GetComment(id int64) (*data.Comment, error) {
	comment, ok := p.comments[id]
	if !ok {
		return nil, fmt.Errorf("comment with ID %d not found", id)
	}
	return comment, nil
}

func (p *TestCommentProvider) UpdateComment(comment *data.Comment) error {
	existingComment, ok := p.comments[comment.Id]
	if !ok {
		return fmt.Errorf("comment with ID %d not found", comment.Id)
	}
	if comment.Content != "" {
		existingComment.Content = comment.Content
	}
	if comment.AuthorId != 0 {
		existingComment.AuthorId = comment.AuthorId
	}
	if comment.RecipeId != 0 {
		existingComment.RecipeId = comment.RecipeId
	}
	if comment.ParentId != 0 {
		existingComment.ParentId = comment.ParentId
	}
	existingComment.UpdatedAt = time.Now()
	return nil
}

func (p *TestCommentProvider) ListRecipeComments(recipeId int64) ([]data.Comment, error) {
	var recipeComments []data.Comment
	for _, comment := range p.comments {
		if comment.RecipeId == recipeId {
			recipeComments = append(recipeComments, *comment)
		}
	}
	return recipeComments, nil
}

func (p *TestCommentProvider) ListUserComments(userId int64) ([]data.Comment, error) {
	var userComments []data.Comment
	for _, comment := range p.comments {
		if comment.AuthorId == userId {
			userComments = append(userComments, *comment)
		}
	}
	return userComments, nil
}

func (p *TestCommentProvider) DeleteComment(id int64) error {
	if _, ok := p.comments[id]; !ok {
		return fmt.Errorf("comment with ID %d not found", id)
	}
	delete(p.comments, id)
	return nil
}

func (p *TestCommentProvider) GetCommentVotes(commentId int64) ([]*data.CommentVote, error) {
	if _, ok := p.commentVotes[commentId]; !ok {
		return nil, fmt.Errorf("comment with ID %d not found", commentId)
	}

	votes := make([]*data.CommentVote, 0)
	for _, commentVotes := range p.commentVotes[commentId] {
		votes = append(votes, commentVotes)
	}

	return votes, nil
}

func (p *TestCommentProvider) AddCommentVote(commentId int64, userId int64, isUpVote bool) error {
	if _, ok := p.commentVotes[commentId]; !ok {
		p.commentVotes[commentId] = make(map[int64]*data.CommentVote)
	}
	p.commentVotes[commentId][userId] = &data.CommentVote{
		Id:        p.nextID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CommentId: commentId,
		UserId:    userId,
		IsUpVote:  isUpVote,
	}
	p.nextID++
	return nil
}

func (p *TestCommentProvider) DeleteCommentVote(commentId int64, userId int64) error {
	if _, ok := p.commentVotes[commentId]; !ok {
		return fmt.Errorf("comment with ID %d not found", commentId)
	}
	delete(p.commentVotes[commentId], userId)
	return nil
}

func (p *TestCommentProvider) GetCommentScore(commentId int64) (int64, error) {
	score := int64(0)
	for _, vote := range p.commentVotes[commentId] {
		if vote.IsUpVote {
			score++
		} else {
			score--
		}
	}
	return score, nil
}

func (p *TestCommentProvider) TearDown() {
	p.nextID = 1
	p.comments = make(map[int64]*data.Comment)
}

func TestCommentService_CreateComment(t *testing.T) {
	provider := NewTestCommentProvider()
	service := NewCommentService(provider)

	type args struct {
		content  string
		authorId int64
		recipeId int64
		parentId int64
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestCommentProvider)
		expect  *data.Comment
	}{
		{
			name: "create top-level",
			args: args{
				content:  "test content",
				authorId: 1,
				recipeId: 1,
			},
			wantErr: false,
			expect: &data.Comment{
				Id:       1,
				Content:  "test content",
				AuthorId: 1,
				RecipeId: 1,
			},
		},
		{
			name: "create nested",
			args: args{
				content:  "test content",
				authorId: 1,
				recipeId: 1,
				parentId: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestCommentProvider) {
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
				})
			},
			expect: &data.Comment{
				Id:       2,
				Content:  "test content",
				AuthorId: 1,
				RecipeId: 1,
				ParentId: 1,
			},
		},
		{
			name: "create multiple nested",
			args: args{
				content:  "test content",
				authorId: 1,
				recipeId: 1,
				parentId: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestCommentProvider) {
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
				})
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
					ParentId: 1,
				})
			},
			expect: &data.Comment{
				Id:       3,
				Content:  "test content",
				AuthorId: 1,
				RecipeId: 1,
				ParentId: 1,
			},
		},
		{
			name: "create multiple nested with same parent",
			args: args{
				content:  "test content",
				authorId: 1,
				recipeId: 1,
				parentId: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestCommentProvider) {
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
				})
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
					ParentId: 1,
				})
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
					ParentId: 1,
				})
			},
			expect: &data.Comment{
				Id:       4,
				Content:  "test content",
				AuthorId: 1,
				RecipeId: 1,
				ParentId: 1,
			},
		},
		{
			name: "err when create nested without existing parent",
			args: args{
				content:  "test content",
				authorId: 1,
				recipeId: 1,
				parentId: 6,
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "err when authorId is not set",
			args: args{
				content:  "test content",
				authorId: 0,
				recipeId: 1,
			},
			wantErr: true,
			expect:  nil,
		},
		{
			name: "err when recipeId is not set",
			args: args{
				content:  "test content",
				authorId: 1,
				recipeId: 0,
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, tc := range testCases {
		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			comment, err := service.CreateComment(tc.args.content, tc.args.authorId, tc.args.recipeId, tc.args.parentId)
			if (err != nil) != tc.wantErr {
				t.Errorf("CreateComment() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !comment.Equal(tc.expect) {
				t.Errorf("CreateComment() = %v, want %v", comment, tc.expect)
			}
		})
		provider.TearDown()
	}
}

func TestCommentService_GetComment(t *testing.T) {
	provider := NewTestCommentProvider()
	service := NewCommentService(provider)

	type args struct {
		id int64
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestCommentProvider)
		expect  *data.Comment
	}{
		{
			name: "get existing comment",
			args: args{
				id: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestCommentProvider) {
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
				})
			},
			expect: &data.Comment{
				Id:       1,
				Content:  "test content",
				AuthorId: 1,
				RecipeId: 1,
			},
		},
		{
			name: "get non-existing comment",
			args: args{
				id: 2,
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, tc := range testCases {
		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			comment, err := service.GetComment(tc.args.id)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetComment() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !comment.Equal(tc.expect) {
				t.Errorf("GetComment() = %v, want %v", comment, tc.expect)
			}
		})
		provider.TearDown()
	}
}

func TestCommentService_UpdateComment(t *testing.T) {
	provider := NewTestCommentProvider()
	service := NewCommentService(provider)

	type args struct {
		id      int64
		content string
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestCommentProvider)
		expect  *data.Comment
	}{
		{
			name: "update existing comment",
			args: args{
				id:      1,
				content: "new content",
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestCommentProvider) {
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
				})
			},
			expect: &data.Comment{
				Id:       1,
				Content:  "new content",
				AuthorId: 1,
				RecipeId: 1,
			},
		},
		{
			name: "update non-existing comment",
			args: args{
				id:      2,
				content: "new content",
			},
			wantErr: true,
			expect:  nil,
		},
	}

	for _, tc := range testCases {
		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			comment, err := service.UpdateComment(tc.args.id, tc.args.content)
			if (err != nil) != tc.wantErr {
				t.Errorf("UpdateComment() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !comment.Equal(tc.expect) {
				t.Errorf("UpdateComment() = %v, want %v", comment, tc.expect)
			}
		})
		provider.TearDown()
	}
}

func TestCommentService_GetCommentScore(t *testing.T) {
	type args struct {
		commentId int64
	}

	testCases := []struct {
		name    string
		args    args
		wantErr bool
		preRun  func(t *testing.T, provider *TestCommentProvider)
		expect  int64
	}{
		{
			name: "get existing comment score",
			args: args{
				commentId: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestCommentProvider) {
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
				})
				provider.AddCommentVote(1, 1, true)
				provider.AddCommentVote(1, 2, false)
				provider.AddCommentVote(1, 3, true)
				provider.AddCommentVote(1, 4, false)
				provider.AddCommentVote(1, 5, true)
				provider.AddCommentVote(1, 6, true)
			},
			expect: 2,
		},
		{
			name: "get non-existing comment score",
			args: args{
				commentId: 2,
			},
			wantErr: true,
			expect:  0,
		},
		{
			name: "get unrated comment score",
			args: args{
				commentId: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestCommentProvider) {
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
				})
			},
			expect: 0,
		},
		{
			name: "get score with multiple downvotes",
			args: args{
				commentId: 1,
			},
			wantErr: false,
			preRun: func(t *testing.T, provider *TestCommentProvider) {
				provider.PutComment(&data.Comment{
					Content:  "test content",
					AuthorId: 1,
					RecipeId: 1,
				})
				provider.AddCommentVote(1, 1, false)
				provider.AddCommentVote(1, 2, false)
				provider.AddCommentVote(1, 3, false)
			},
			expect: -3,
		},
	}

	for _, tc := range testCases {
		provider := NewTestCommentProvider()
		service := NewCommentService(provider)

		if tc.preRun != nil {
			tc.preRun(t, provider)
		}
		t.Run(tc.name, func(t *testing.T) {
			score, err := service.GetCommentScore(tc.args.commentId)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetCommentScore() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if score != tc.expect {
				t.Errorf("GetCommentScore() = %v, want %v", score, tc.expect)
			}
		})
		provider.TearDown()
	}
}
