defaultMessage="deafult message from git script"
defaultBranch=master

git add .
git commit -m "${1:-$defaultMessage}"
git push origin "${2:-$defaultBranch}"