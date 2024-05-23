# rss-feed-aggregator

## How to use

### Setup PostgreSQL database

- Make sure your postgres database is running and have set up goose and sqlc
- Run migration
  - `cd sql/schema && goose postgres postgres://postgres:postgres@localhost:5432/dbname up`
  - (not required) run `sqlc generate` to generate go code from SQL
- Run server: `go build && ./rss`
  The server will run on port `8080`.

## API documentation

### Overview

RSS Feed Aggregator API provides endpoints to retrieve user subscribed feeds and posts.

#### Check the status of the service:

- Method: `GET`
- URL: `http://localhost:8080/readiness`
- Success Response:
  - Code: 200
  - Body:
    ```json
    { "status": "ok" }
    ```

#### Authentication

Users table contains `API_KEY` field which is used to authenticate requests.

#### Create a user:

- Method: `POST`
- URL: `http://localhost:8080/v1/users`
- Body: `{"name": "YOUR_USERNAME"}`

- Success Response:

  - Code: 201
  - Body:
    ```json
    {
      "created_at": "2024-05-23T15:51:44.362978Z",
      "updated_at": "2024-05-23T15:51:44.362995Z",
      "api_key": "38393544fa6700e12f990e1a10459c67e5b89c080e00799fef6253b197e58e7e",
      "name": "YOUR_USERNAME",
      "id": "e310ece5-a3fa-40c7-9259-4e5fb21a5c8f"
    }
    ```

#### Get a user:

- Method: `GET`
- URL: `http://localhost:8080/v1/users/`
- Header: `{"ApiKey": "YOUR_API_KEY"}`
- Success Response:
  - Code: 200
  - Body:
    ```json
    {
      "created_at": "2024-05-23T15:51:44.362978Z",
      "updated_at": "2024-05-23T15:51:44.362995Z",
      "api_key": "38393544fa6700e12f990e1a10459c67e5b89c080e00799fef6253b197e58e7e",
      "name": "YOUR_USERNAME",
      "id": "e310ece5-a3fa-40c7-9259-4e5fb21a5c8f"
    }
    ```

#### Create Feeds:

Feed is the link to the blog you want to subscribe. Right now only RSS links are supported.
Feed Follow contains the user id and feed id to which user is subscribed to. When a user creates a feed, they automatically get subscribed to that feed.

- Method: `POST`
- URL: `http://localhost:8080/v1/feeds`
- Header: `{"ApiKey": "YOUR_API_KEY"}`
- Body:
  ```json
  {
    "name": "YOUR_FEED_NAME",
    "url": "YOUR_XML_RSS_FEED_LINK"
  }
  ```
- Success Response:
  ```json
  {
    "feed": {
      "created_at": "2024-05-23T13:21:41.091793Z",
      "updated_at": "2024-05-23T13:21:41.091793Z",
      "last_fetched_at": "0001-01-01T00:00:00Z",
      "url": "YOUR_XML_RSS_FEED_LINK",
      "name": "YOUR_FEED_NAME",
      "user_id": "127cc657-adf2-4661-9321-dbaa5341c77b",
      "id": "6b409740-011f-4f86-bfe4-ceb2db9fc52f"
    },
    "feed_follow": {
      "created_at": "2024-05-23T13:21:41.095761Z",
      "updated_at": "2024-05-23T13:21:41.095761Z",
      "id": "e4ddc178-62a2-4a9c-a94b-19003e6e8a62",
      "feed_id": "6b409740-011f-4f86-bfe4-ceb2db9fc52f",
      "user_id": "127cc657-adf2-4661-9321-dbaa5341c77b"
    }
  }
  ```
- Error Response :
  - Code: 409
  ```json
  {
    "error": "Feed already exists"
  }
  ```

#### Get All Feeds:

This endpoint get all the fields in the database.

- Method: `GET`
- URL: `http://localhost:8080/v1/feeds`
- Header: `{"ApiKey": "YOUR_API_KEY"}`
- Success Response:
  ```json
  {
    [
      {
        "created_at": "2024-05-23T13:21:41.091793Z",
        "updated_at": "2024-05-23T13:21:41.091793Z",
        "last_fetched_at": "0001-01-01T00:00:00Z",
        "url": "YOUR_XML_RSS_FEED_LINK",
        "name": "YOUR_FEED_NAME",
        "user_id": "127cc657-adf2-4661-9321-dbaa5341c77b",
        "id": "6b409740-011f-4f86-bfe4-ceb2db9fc52f"
      },
      {
       "created_at": "2024-05-23T13:21:41.091793Z",
        "updated_at": "2024-05-23T13:21:41.091793Z",
        "last_fetched_at": "0001-01-01T00:00:00Z",
        "url": "YOUR_XML_RSS_FEED_LINK2",
        "name": "YOUR_FEED_NAME2",
        "user_id": "127cc657-adf2-4661-9321-dbaa5341c77b",
        "id": "6b409740-011f-4f86-bfe4-ceb2db9fc52f"
      }
    ]
  }
  ```

#### Follow a Feed:

To subscribe to a feed a user can use this endpoint, it requires user api key and feed id.

- Method: `POST`
- URL: `http://localhost:8080/v1/feed_follows/`
- Header: `{"ApiKey": "YOUR_API_KEY"}`
- Body: `{"feed_id": "YOUR_FEED_ID"}`
- Success Response:

  - Code: 201
  - Body:

  ```json
  {
    "created_at": "2024-05-23T19:26:46.830128Z",
    "updated_at": "2024-05-23T19:26:46.830127Z",
    "id": "623d8856-fe28-4144-8eb4-782e268685c7",
    "feed_id": "22cb86da-0d63-4aed-b081-f70a9db6ff7d",
    "user_id": "254b3729-c947-47b0-b80a-401ebd972f65"
  }
  ```

#### Unfollow a Feed:

To unsubscribe to a feed a user can use this endpoint, it requires user API key and feed follow id.

- Method: `DELETE`
- URL: `http://localhost:8080/v1/feed_follows/feedFollowId=YOUR_FEED_FOLLOW_ID`
- Header: `{"ApiKey": "YOUR_API_KEY"}`
- Success Response:
  - Code: 201
  - Body: `{"message": "Feed follow deleted"}`
- Error Response:
  - Invalid feed follow id
  - Feed follow not found
  - Internal server error

#### Get All Feed Follows:

This get all the feeds that a user is subscribed to.

- Method: GET
- URL: `http://localhost:8080/v1/feed_follows`
- Header: `{"ApiKey": "YOUR_API_KEY"}`
- Success Response:
  - Code 200
  - Body:
  ```json
  [
    {
      "created_at": "2024-05-23T00:55:23.545983Z",
      "updated_at": "2024-05-23T00:55:23.545983Z",
      "id": "4b4a697f-e16b-4db0-9e6b-8ddb7fe6eeb0",
      "feed_id": "bc1bf3c3-af3d-4236-9158-f464392cb81d",
      "user_id": "0ca6d58d-05d6-47e6-bf74-384f7afb5799"
    },
    {
      "created_at": "2024-05-23T00:55:04.478651Z",
      "updated_at": "2024-05-23T00:55:04.478651Z",
      "id": "ef135f5e-3c35-425a-a126-662b6210d028",
      "feed_id": "b1124b6d-7dce-4690-a10f-f050af2f8e11",
      "user_id": "0ca6d58d-05d6-47e6-bf74-384f7afb5799"
    },
    {
      "created_at": "2024-05-22T13:04:12.044244Z",
      "updated_at": "2024-05-22T13:04:12.044244Z",
      "id": "5290968c-d14a-4fee-8058-449339f5f9e3",
      "feed_id": "a98f4ef2-d397-4927-aa84-b817977cdbbf",
      "user_id": "0ca6d58d-05d6-47e6-bf74-384f7afb5799"
    }
  ]
  ```

#### Get Posts for a user:

Posts are all the main contents (title and description) of the RSS feed. These post are updated every specified duration of time. The limit is the amount of posts you want to receive in a go.

- Method: GET
- URL: `http://localhost:8080/v1/posts?limit=10/`
- Header: `{"ApiKey": "YOUR_API_KEY"}`
- Success Response:
  - Code 200
  - Body:
  ```json
  [
    {
      "created_at": "2024-05-22T07:28:19.705485Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "title": "Feed: All Latest",
      "url": "https://www.wired.com/feed",
      "description": "Channel Description",
      "published_at": "",
      "id": "5bcdcef2-1045-4523-aaaa-b59bddaaffde",
      "feed_id": "a98f4ef2-d397-4927-aa84-b817977cdbbf"
    },
    {
      "created_at": "2024-05-22T07:28:19.90791Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "title": "NPR Topics: News",
      "url": "http://www.npr.org/rss/rss.php?id=1001",
      "description": "NPR news, audio, and podcasts. Coverage of breaking stories, national and world news, politics, business, science, technology, and extended coverage of major national and world events.",
      "published_at": "",
      "id": "0655eae9-e9a4-404a-99f9-6babdb335823",
      "feed_id": "bb66774e-4c46-4af6-9111-685c7189af79"
    },
    {
      "created_at": "2024-05-22T07:28:19.765071Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "title": "The Hacker News",
      "url": "https://feeds.feedburner.com/TheHackersNews",
      "description": "Most trusted, widely-read independent cybersecurity news source for everyone; supported by hackers and IT professionals — Send TIPs to admin@thehackernews.com",
      "published_at": "",
      "id": "44fabba0-1d44-42c5-9d63-2ad508883d79",
      "feed_id": "2f6aed12-f123-4163-98e8-5a620bf5e3c4"
    },
    {
      "created_at": "2024-05-22T07:28:20.100606Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "title": "Washington Post",
      "url": "https://feeds.washingtonpost.com/rss/world?itid=lk_inline_manual_35",
      "description": "Washington Post News Feed",
      "published_at": "",
      "id": "1205ef70-2ed8-459b-be65-10824832c67d",
      "feed_id": "8d3a4f6c-f868-4970-a537-b67352bcf151"
    }
  ]
  ```

#### Get All the RSS items of all the feeds user is subscribed to:

RSS items are the contents inside the RSS feed. These are also updated along with the RSS posts.
These contains unique stories/ articles and there link. The database stores all the past stories for a user.

- Method: GET
- Url: `http://localhost:8080/v1/items`
- Header: `{"ApiKey": "YOUR_API_KEY"}`
- Success Response:
  - Code: 200
  - Body:
  ```json
  [
    {
      "id": "160c741e-8735-446c-9146-4f154c5cf17e",
      "created_at": "2024-05-22T07:28:19.709574Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "title": "/e/OS Review: This Operating System Is Better Than Android. You Should Try It",
      "url": "https://www.wired.com/story/e-os-review/",
      "author": "",
      "description": "Even if you aren’t technically inclined, this privacy-first mobile operating system helps you escape surveillance capitalism.",
      "published_at": "Wed, 22 May 2024 07:00:00 +0000",
      "post_id": "5bcdcef2-1045-4523-aaaa-b59bddaaffde"
    },
    {
      "id": "02cd5d7f-274a-4b2a-a778-e5bd3b8855f7",
      "created_at": "2024-05-22T07:28:19.71298Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "title": "Neuralink’s First User Is ‘Constantly Multitasking’ With His Brain Implant",
      "url": "https://www.wired.com/story/neuralink-first-patient-interview-noland-arbaugh-elon-musk/",
      "author": "",
      "description": "Noland Arbaugh is the first to get Elon Musk’s brain device. The 30-year-old speaks to WIRED about what it’s like to use a computer with his mind—and gain a new sense of independence.",
      "published_at": "Wed, 22 May 2024 06:00:00 +0000",
      "post_id": "5bcdcef2-1045-4523-aaaa-b59bddaaffde"
    },
    {
      "id": "7b257a96-65ee-418e-a506-950ecd68a680",
      "created_at": "2024-05-22T07:28:19.715902Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "title": "A Far-Right Indian News Site Posts Racist Conspiracies. US Tech Companies Keep Platforming It",
      "url": "https://www.wired.com/story/india-opindia-google-facebook-advertising/",
      "author": "",
      "description": "OpIndia claims “Islamophobia does not exist.” A new report shared exclusively with WIRED finds Google’s programmatic ads are running next to its content.",
      "published_at": "Wed, 22 May 2024 06:00:00 +0000",
      "post_id": "5bcdcef2-1045-4523-aaaa-b59bddaaffde"
    },
    {
      "id": "2417bd78-753d-4405-a2d2-ec1589e535c2",
      "created_at": "2024-05-22T07:28:19.719169Z",
      "updated_at": "0001-01-01T00:00:00Z",
      "title": "The 36 Best Shows on Hulu Right Now (May 2024)",
      "url": "https://www.wired.com/story/best-tv-shows-hulu-this-week/",
      "author": "",
      "description": "Black Twitter: A People’s History, Shōgun, and Under the Bridge are just a few of the shows you should be watching on Hulu this month.",
      "published_at": "Tue, 21 May 2024 19:00:00 +0000",
      "post_id": "5bcdcef2-1045-4523-aaaa-b59bddaaffde"
    },
  .... and so on
  ]
  ```
