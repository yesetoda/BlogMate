# BlogMate

BlogMate is a powerful, feature-rich blogging API written in Go. Built on the Gin framework and backed by MongoDB, BlogMate supports everything you need for a modern blog platform—including user management, blog post creation, advanced comment and reply interactions, and AI-powered content recommendations. Our system integrates with Gemini AI to provide tailored suggestions for blog titles, content, and tags, ensuring your content meets your validation rules and adheres to your standards.

BlogMate is maintained by the [Yesetoda](https://github.com/yesetoda) team and is designed for seamless integration into your development pipeline.

---

## Table of Contents

- [Features](#features)
- [Architecture & Tech Stack](#architecture--tech-stack)
- [Installation & Setup](#installation--setup)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [AI Model & Validation](#ai-model--validation)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

---

## Features

- **User Management:**  
  Secure registration, login, account verification, token refresh, and role-based authorization using JWT.

- **Blog Management:**  
  Create, update, delete, and retrieve blog posts with complete CRUD support, including pagination and filtering.

- **Interactive Comments & Replies:**  
  Post comments and replies on blog posts with capabilities for likes, dislikes, and views.

- **AI-Powered Recommendations:**  
  Leverage Gemini AI to generate content suggestions including:
  - **Blog Recommendations:**  
    ```go
    RecommendBlogs(Data) ([]domain.BlogRecommendation, error)
    ```
  - **Title Suggestions:**  
    ```go
    RecommendTitle(content string, tags []string) (string, error)
    ```
  - **Content Generation:**  
    ```go
    RecommendContent(title string, tags []string) (string, error)
    ```
  - **Tag Recommendations:**  
    ```go
    RecommendTags(title string, content string) ([]string, error)
    ```
  - **Content Summarization:**  
    ```go
    Summarize(Data) (string, error)
    ```
  - **Validation:**  
    ```go
    Validate(Data) error
    ```
  - **Chat Interface:**  
    ```go
    Chat(prompt string) (string, error)
    ```

- **Email Support:**  
  BlogMate supports emailing capabilities via Google App credentials for sending verification emails and other notifications.

- **Validation & Rule Enforcement:**  
  With integrated validation, BlogMate ensures that all blog content adheres to defined rules before it is published.

- **Swagger Documentation:**  
  Comprehensive API documentation using Swaggo is built-in. Easily explore and test endpoints with the interactive Swagger UI.

---

## Architecture & Tech Stack

- **Language:** Go (Golang)
- **Web Framework:** Gin
- **Database:** MongoDB
- **Documentation:** Swagger (using [Swaggo](https://github.com/swaggo/swag))
- **AI Integration:** Gemini AI for content recommendation (via configurable API keys and models)
- **Email Integration:** Google App credentials for email notifications
- **Configuration:** Custom configuration via a YAML file (no reliance on .env)

---

## Installation & Setup

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/yesetoda/BlogMate.git
   cd BlogMate
   ```

2. **Install Dependencies:**

   Ensure you have [Go modules](https://github.com/golang/go/wiki/Modules) enabled:

   ```bash
   go mod download
   ```

3. **Build the Application:**

   Build the project to verify that everything is set up properly:

   ```bash
   cd delivery
   go build -o blogmate
   ```

---

## Configuration

BlogMate uses a YAML configuration file to define all the settings for the application. Create a configuration file (for example, `config.yaml`) with the following structure:

```yaml
database:
  username: <your-database-username>
  password: <your-database-password>
  uri: <your-database-uri>
email:
  key: <google app key>
port: <port number>
jwt: <jwt key>
gemini:
  api_key: <your-gemini-key>
  model: <gemini-model-name>
```

Using this approach, you can easily manage environment-specific settings without relying on `.env` files.

---

## Running the Application

Once your configuration file is ready and your dependencies are installed, you can run BlogMate with:

```bash
cd delivery
go run main.go
```

BlogMate will start the server on the port specified in your configuration. You will see log messages confirming the startup and available endpoints.

---

## API Documentation

BlogMate comes with built-in interactive API documentation using Swagger. To generate the Swagger docs, run:

```bash
swag init --dir . --parseDependency
```

After generation, start the server and navigate to:

```
http://localhost:<port>/docs/index.html
```

This interface lets you interact with all API endpoints, view detailed information on request/response models, and even authorize endpoints using JWT Bearer tokens.

---

## AI Model & Validation

The AI model interface provides content recommendations and validation functions:

```go
type AIModel interface {
	RecommendBlogs(Data) ([]domain.BlogRecommendation, error)
	RecommendTitle(content string, tags []string) (string, error)
	RecommendContent(title string, tags []string) (string, error)
	RecommendTags(title string, content string) ([]string, error)
	Summarize(Data) (string, error)
	Validate(Data) error
	Chat(prompt string) (string, error)
}
```

- **Recommendations:** Generate suggestions for blog titles, content, tags, and complete blog recommendations.
- **Validation:** Check that the blog content adheres to defined rules.
- **Chat:** A conversational interface to interact with the AI.

To integrate and configure Gemini AI, ensure your configuration file has the correct `gemini.api_key` and `gemini.model` values. The AI service is initialized when adding the AI routes, and errors will be logged if the API key is missing or invalid.

---

## Contributing

Contributions to BlogMate are always welcome. If you’d like to contribute:

1. Fork the repository on GitHub.
2. Create a new branch with your feature or bug fix.
3. Submit a pull request with a detailed description of your changes.

Make sure to follow the repository’s coding standards and include tests where appropriate.

---

## License

BlogMate is open-source software.

---

## Contact

For questions, support, or feedback, please contact the Yesetoda team at:  
**Email:** yeneinehseiba@gmail.com

Or visit our [GitHub repository](https://github.com/yesetoda/BlogMate) for more details.
