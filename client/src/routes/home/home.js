import React from 'react'
import {gql} from 'apollo-boost';
import {compose, graphql} from 'react-apollo';

// example of a query
const query = gql`
  query GetGuidelines{
    getGuidelines {
      title
      description
    }
  }
`;

// example of a graphql mutation
const mutation = gql`
  mutation CreateGuideline($title: String, $description: String) {
    createGuideline(title: $title, description: $description) {
      title
    }
  }
`;

// queries run automatically!
@compose(
    graphql(query),
    graphql(mutation)
)
class Home extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            title: '',
            description: ''
        }
    }

    onChange = event => {
        this.setState({[event.target.name]: event.target.value})
    };

    save = () => {
        this.props.mutate({
            variables: {
                title: this.state.title,
                description: this.state.description
            },
            refetchQueries: ['GetGuidelines']
        })
    };

    render() {
        const {getGuidelines = []} = this.props.data;
        return (

            <div style={{display: 'inline', justifyContent: 'center', marginTop: '5%'}}>
                <div style={{display: 'inline', flexDirection: 'center'}}>
                    <h3 style={{margin: '0px'}}>Add new guideline!</h3>
                    <input placeholder='Title' name='title' onChange={this.onChange}/><br/>
                    <input placeholder='Description' name='description' onChange={this.onChange}/><br/>
                    <button onClick={this.save}>Save</button>
                    <br/>
                </div>
                <br/>

                <table style={{border: '1px solid black'}}>
                    <tbody>
                    <tr>
                        <th style={{border: '1px solid black'}}>index</th>
                        <th style={{border: '1px solid black'}}>title</th>
                        <th style={{border: '1px solid black'}}>description</th>
                    </tr>
                    {
                        getGuidelines.map((notTodo, index) => (
                            <tr key={index}>
                                <td style={{border: '1px solid black'}}>{index}</td>
                                <td style={{border: '1px solid black'}}>{notTodo.title}</td>
                                <td style={{border: '1px solid black'}}>{notTodo.description}</td>
                            </tr>
                        ))
                    }
                    </tbody>
                </table>

            </div>
        )
    }
}

export default Home