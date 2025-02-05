import requests

def get_pull_requests(github_base_url :str, user_name :str, repository :dict, headers :dict) -> tuple[dict, str]:
    github_pull_requests_url = f"{github_base_url}/repos/{user_name}/{repository['name']}/pulls"
    response = requests.request("GET", github_pull_requests_url, headers=headers)

    if response.status_code == 200:
        pull_requests = response.json()
        return pull_requests, ""
    else:
        return {}, f"Failed to retrieve pull requests. Status code: {response.status_code}"
