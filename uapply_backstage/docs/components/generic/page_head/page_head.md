# Page Head 

## Where to use ?

After login , in the super master and master Page , right top.

## For What ?

Three element : 
- title of the login organization
- welcome depending on your role 
- sign out

## Path 

> src/components/generic/page_head/page_head.vue

## Props

| Name | Type | Instruction |
|------|------|-------------|
| title |  string  | the title of the organization |
| role |  string  | the role of the user |

## Emits

| Name | Props | Instruction | Trigger |
|------|------|--------------|---------|

## Defined Variable

| Name | Type | Instruction | Required |
|------|------|-------------|----------|
| user_title | string | the welcome centence | true |

## Events

| Name | Instruction | Callback Parameter |
|------|-------------|--------------------|
| sign_out | sign out this account and go to login page | / |


## Need Adaptation

| Where | What | Supplement | 
|-------|------|------------|
| right part | the role title and button maybe conflict | / |