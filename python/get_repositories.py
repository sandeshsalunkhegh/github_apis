import requests

def get_repositories(github_base_url :str, user_name :str, headers :dict) -> tuple[dict, str]:
    github_repositories_url = f"{github_base_url}/users/{user_name}/repos"
    response = requests.request("GET", github_repositories_url, headers=headers)

    if response.status_code == 200:
        repositories = response.json()
        return repositories, ""
    else:
        return {}, f"Failed to retrieve repositories. Status code: {response.status_code}"
