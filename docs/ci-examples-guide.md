# CI Examples Guide

## GitHub Actions

### Step that gets environment variables from github to be used in the CI process and that will be used in the version increment strategy
````
extract-gha-vars:
    name: Extract GHA Variables
    #needs: [sonarqube]
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v2
      with:
        fetch-depth: 0

    # extract branch name
    - name: Extract branch name
      if: github.event_name != 'pull_request'
      id: branch_name
      shell: bash
      run: | 
        echo "Branch Name: ${GITHUB_REF#refs/heads/}"
        echo "::set-output name=branch_name::$(echo ${GITHUB_REF#refs/heads/})"

    # extract commit message
    - name: Extract commit message
      if: github.event_name != 'pull_request'
      id: commit_msg
      shell: bash
      run: echo "::set-output name=cmt_msg::$(echo ${{ github.event.head_commit.message }})"
    
    # extract commit message
    - name: Extract commit type
      if: github.event_name != 'pull_request'
      id: commit_type
      shell: bash
      run: echo "::set-output name=cmt_type::$(echo ${{ github.event.head_commit.message }} | awk -F":" '{print $1}')"
    
    - name: Extract Environment Deploy
      id: env_deploy
      run: |
        env_deploy=`echo ${{ github.event.head_commit.message }} | awk -F":" '{print $2}'`
        if [[ ($env_deploy == "") || 
              ($env_deploy != "dev") || 
              ($env_deploy != "test") || 
              ($env_deploy != "prod") || 
              ($env_deploy != "uat") ]];then env_deploy=${{ env.DEFAULT_ENV_GERNERATE_TAG_RELEASE }};fi

        env_deploy_dispatch=${{ github.event.inputs.env_deploy_dispatch }}
        if [[ $env_deploy_dispatch != "" ]];then env_deploy=$env_deploy_dispatch;fi

        echo "GCP Project: ${{ env.GCP_PROJT }}"
        echo "Environment: $env_deploy"
        echo "Flag DEFAULT_ENV_GERNERATE_TAG_RELEASE: ${{ env.DEFAULT_ENV_GERNERATE_TAG_RELEASE }}"

        echo "::set-output name=env_deploy::$(echo $env_deploy)"
        echo "::set-output name=def_env_gen_tag_rel::$(echo ${{ env.DEFAULT_ENV_GERNERATE_TAG_RELEASE }})"

    - name: Extract Last GitHub Tag Version
      id: github_tag_version
      run: |
        github_tag_version=$(git tag -l --sort=-v:refname | head -1)
        if [[ $github_tag_version == ""  ]];then github_tag_version="0.0.0"; fi

        echo "GitHub Tag Version: $github_tag_version"
        echo "::set-output name=gthub_tag_version::$(echo $github_tag_version)"
        
    outputs:
      branch_name: ${{ steps.branch_name.outputs.branch_name }}
      commit_msg:  ${{ steps.commit_msg.outputs.cmt_msg }}
      cmt_type:  ${{ steps.commit_type.outputs.cmt_type }}
      env_deploy: ${{ steps.env_deploy.outputs.env_deploy }}
      is_rollback: ${{ steps.env_deploy.outputs.is_rollback }}
      default_env_generate_tag_release: ${{ steps.env_deploy.outputs.def_env_gen_tag_rel }}
      github_tag_version: ${{ steps.github_tag_version.outputs.gthub_tag_version }}

````

### versctl checkout step (Use a token that has permission to access the repository)
````
 - name: Checkout
      uses: actions/checkout@v2
      with:
        token: ${{ secrets.TOKEN }}
        repository: 'jonascavalcanti/versctl'
        path: 'versctl'
````

### Increment step (This step uses commit conventions and semantic version specification)
```
- name: Increment application version
      id: increment
      env:
        versctl_bin: "./versctl/versctl-linux-amd64"
      shell: bash
      run: |
        set -x
        semver="patch"
        env_deploy=`echo ${{ needs.extract-gha-vars.outputs.env_deploy }}`
        cmt_type=`echo ${{ needs.extract-gha-vars.outputs.cmt_type }}`
        if [[ $cmt_type == "feat" ]];then semver="minor";fi
        if [[ $cmt_type == "BREAKING CHANGE" ]];then semver="major";fi
        
        github_tag_version=`echo ${{ needs.extract-gha-vars.outputs.github_tag_version }}`
        
        chmod +x $versctl_bin
        app_version=`$versctl_bin update --version $github_tag_version -i $semver`

        echo "Semantic Versioning: $semver"
        echo "Application Version: $github_tag_version"
        echo "Application New Version: $app_version"

        echo "::set-output name=app_version::$(echo $app_version)"
    outputs:
      app_version: ${{ steps.increment.outputs.app_version }}
```
### Get the app_version variable inside CI and send to another steps
````
${{ needs.generate-version.outputs.app_version }}
````

