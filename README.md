Boilerplate for creating an API using clean architecture.

Using :

- Go
- fiber (v2)
- gorm (PSQL Driver)
- jwt
- validator
- uuid
- godotenv
- logrus
- fiber/swagger
- swaggo/swag

Todo :

- Redis cache
- Websocket on game start
- AWS S3 for images

# Inspirations

- https://github.com/RizkiMufrizal/gofiber-clean-architecture
- https://github.com/Creatly/creatly-backend/
-

# API

This API manage a quizz app (the frontend repo will come later).

- Users :
  - A user can signup, login, update his informations, create questions and create/join a game
  - A user can have a role : administrator, moderator, or user
  - During a game, user can upvote or downvote a question
- Game :
  - A game has a list of parameters (number of questions, no downvoted questions, ...)
  - Has a list of participating user
  - A game is created by a user, saved in memory during game, and saved in database after
  - A user can ask for his games history
- Question :
  - A question has a label, a good answer, and a minimum of 1 (and max 10) wrong answers, has a field CreatedBy (user)
- Vote :
  - Relation between a User and a Question - Type Upvote/Downvote
