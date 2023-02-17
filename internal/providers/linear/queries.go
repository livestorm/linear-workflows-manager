package linear

const getUserQuery = `
	query getUser($identifier: String = "userid") {
		user(id: $identifier) {
			email
			name
		}
	}
`

const getStatesQuery = `
	query getStates($team: String = "Tech") {
		workflowStates(filter: {team: {name: {eq: $team}}}) {
			nodes {
				id
				name
			}
		}
	}
`

const changeStateMutation = `
	mutation updateIssue($identifier: String = "SUP-2250", $stateId: String = "") {
		issueUpdate(id: $identifier, input: {stateId: $stateId}) {
			success
		}
	}
`

const addCommentMutation = `
	mutation addComment($identifier: String = "SUP-2250", $comment: String = ""){
		commentCreate(input: { issueId: $identifier, body: $comment}) {
			success
		}
  	}
`
