import React from "react"
import ReactDOM from "react-dom"
import styled from "styled-components"
import {postJson, colorScale, appdata} from "./constants/AppConstants"
import {BrowserRouter as Router, Link, Route} from "react-router-dom"


class App extends React.Component {
    constructor(props) {
        super(props);

        //this.logout = this.logout.bind(this);
    }

    // logout() {
    //     // localStorage.removeItem("id_token");
    //     // localStorage.removeItem("access_token");
    //     // localStorage.removeItem("profile");
    //     // location.reload();
    // }

    render() {
        return (
            <Router>
                <div className="container">
                    {/*<span className="pull-right">*/}
                    {/*    <a onClick={this.logout}>Log out</a>*/}
                    {/*</span>*/}
                    <Title>
                        Spotify Insights
                    </Title>

                    <Row>
                        <div className="row">
                            <div className="col-sm-4">
                                <Text>
                                    Spotify Visualizations
                                </Text>
                                <Link to={`/genre-count`}>
                                    <button className="btn btn-primary">Genre Count</button>
                                </Link>

                            </div>
                        </div>
                    </Row>

                    <Route path="/genre-count" component={GenreCountContainer} />
                </div>
            </Router>
        );
    }
}

class GenreCountContainer extends React.Component {
    constructor(props) {
        super(props);
        this.serverRequest = this.serverRequest.bind(this);
        this.state = {
            genreCounts: [],
            count: 0,

        };
        
    }

    serverRequest() {
        $.get("http://localhost:3000/api/genre/all/counts", res => {
            this.setState({
                genreCounts: res
            });
        });
    }

    componentDidMount() { // ran after inserted into DOM
        this.serverRequest();
    }

    componentWillUnmount() { // run when component will unmount from dom
    }

    render() {
        return(
                <div className="row">
                    { Object.keys(this.state.genreCounts).map((key, index) => (
                        //<p key={index}> this is my key {key} and this is my value {this.state.genreCounts[key]}</p>
                        <GenreCountCard key={index} genreName={key} count={this.state.genreCounts[key]} />
                    ))}
                </div>
        );
    }
}

class GenreCountCard extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            //name: "",
            count: 0,
        };
        // this.genre = this.genre.bind(this);
        // this.serverRequest = this.serverRequest.bind(this);
    }

    // like() {
    //     let joke = this.props.joke;
    //     joke.likes += 1; // update single joke (persistent data)
    //     this.serverRequest(joke);
    // }

    // async serverRequest(joke) {
    //     const response = await fetch(`http://localhost:3000/api/jokes/like/${joke.id}`, postJson);
    //     const data = await response;
    //     this.setState({ liked: "Liked!", jokes: data});
    //     this.props.jokes = data;
    // }

    render() {
        return (
            <div className="col-sm-4">
                <div className="panel panel-default">
                    <div className="panel-heading"><b>{this.props.genreName}</b></div>
                    <div className="panel-body">{this.props.count}</div>
                    <div className="panel-footer">
                    </div>
                </div>
            </div>
        )
    }
}



const Title = styled.h1`
  text-align: center;
  color: #1DB954;
`;

const Text = styled.p`
  color: #fff;
`;

// const Row = styled.div.attrs({
//     className: 'row-bottom-margin'
// })`
// margin-bottom: 100px;
// `;

const Row = styled.div`
  margin-bottom: 20px;
`

ReactDOM.render(<App />, document.getElementById('app'));