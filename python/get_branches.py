import requests

def get_branches(github_base_url :str, user_name :str, repository :dict, headers :dict) -> tuple[dict, str]:
    github_branches_url = f"{github_base_url}/repos/{user_name}/{repository['name']}/branches"
    response = requests.request("GET", github_branches_url, headers=headers)

    if response.status_code == 200:
        branches = response.json()
        return branches, ""
    else:
        return {}, f"Failed to retrieve branches. Status code: {response.status_code}"
