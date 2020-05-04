endpoint "FetchLatestProject" "http" {
	endpoint = "/"
	method = "GET"
	codec = "json"
}

flow "FetchLatestProject" {
	input "com.maestro.Query" {}

	resource "query" {
		request "com.maestro.Service" "GetTodo" {
		}
	}

	resource "user" {
		request "com.maestro.Service" "GetUser" {
		}
	}

	output "com.maestro.Item" {
		header {
			Username = "{{ user:username }}"
		}

		id = "{{ query:id }}"
		title = "{{ query:title }}"
		completed = "{{ query:completed }}"
	}
}
