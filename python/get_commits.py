import requests

def get_commits(github_base_url :str, user_name :str, repository :dict, headers :dict) -> tuple[dict, str]:
    github_commits_url = f"{github_base_url}/repos/{user_name}/{repository['name']}/commits"
    response = requests.request("GET", github_commits_url, headers=headers)

    if response.status_code == 200:
        commits = response.json()
        return commits, ""
    else:
        return {}, f"Failed to retrieve commits. Status code: {response.status_code}"
