from dotenv import load_dotenv
from get_repositories import get_repositories
from get_pull_requests import get_pull_requests
from get_commits import get_commits
from get_branches import get_branches

import json
import os
import sys

load_dotenv('../secrets.env')

GITHUB_BASE_URL = 'https://api.github.com' 
USER_NAME = 'sandeshsalunkhegh'
GITHUB_ACCESS_TOKEN = os.getenv('GITHUB_ACCESS_TOKEN')

github_repositories_url = f"{GITHUB_BASE_URL}/users/{USER_NAME}/repos"

headers = {
    "Authorization": f"Bearer {GITHUB_ACCESS_TOKEN}",
    "Accept": "application/vnd.github.v3+json",
}

(repositories, error) = get_repositories(GITHUB_BASE_URL, USER_NAME, headers)

if error:
    print(error)
    sys.exit(1)

for repository in repositories:
    print(json.dumps(repository, indent=4))

    (pull_requests, error) = get_pull_requests(GITHUB_BASE_URL, USER_NAME, repository, headers)
    if error:
        print(error)
        sys.exit(1)
    if pull_requests:
        print(f"Pull requests for repository: {repository['name']}")
        print(json.dumps(pull_requests, indent=4))
    else:
        print(f"No pull requests for repository: {repository['name']}")
    
    (commits, error) = get_commits(GITHUB_BASE_URL, USER_NAME, repository, headers)
    if error:
        print(error)
        sys.exit(1)
    if commits:
        print(f"Commits for repository: {repository['name']}")
        print(json.dumps(commits, indent=4))
    else:
        print(f"No commits for repository: {repository['name']}")

    (branches, error) = get_branches(GITHUB_BASE_URL, USER_NAME, repository, headers)
    if error:
        print(error)
        sys.exit(1)
    if branches:
        print(f"Branches for repository: {repository['name']}")
        print(json.dumps(commits, indent=4))
    else:
        print(f"No branches for repository: {repository['name']}")
