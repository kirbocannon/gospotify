import React from "react";
import ReactDOM from "react-dom";
import {postJson} from "../constants/AppConstants"


class App extends React.Component {

    render() {
        this.loggedIn = true;
        if (this.loggedIn) {
            return (<LoggedIn />);
        } else {
            return (<Home />);
        }
    }
}

class Home extends React.Component {
    render() {
        return (
            <div className="container">
                <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
                    <h1>Spotify Insigns</h1>
                    <p>Insight into your personal Spotify Account</p>
                    <p>Sign in to get access </p>
                    <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
                </div>
            </div>
        )
    }
}

class LoggedIn extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            genreCounts: []
        };

        this.serverRequest = this.serverRequest.bind(this);
        this.logout = this.logout.bind(this);
    }

    logout() {
        // localStorage.removeItem("id_token");
        // localStorage.removeItem("access_token");
        // localStorage.removeItem("profile");
        // location.reload();
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

    render() {
        return (
            <div className="container">
                <br />
                <span className="pull-right">
          <a onClick={this.logout}>Log out</a>
        </span>
                <h2>Spotify Insights</h2>
                <p>Spotify Visualizations</p>
                <div className="row">
                    <div className="container">
                        {/*{this.state.genreCounts.map(function(joke, _) {*/}
                        {/*    return <Joke key={joke.id} joke={joke} />;*/}
                        {/*})}*/}
                        {  Object.keys(this.state.genreCounts).map((key, index) => (
                            //<p key={index}> this is my key {key} and this is my value {this.state.genreCounts[key]}</p>
                            <GenreCount key={index} genreName={key} count={this.state.genreCounts[key]} />
                        ))}
                    </div>
                </div>
            </div>
        );
    }
}

class GenreCount extends React.Component {
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
            <div className="col-xs-4">
                <div className="panel panel-default">
                    <div className="panel-heading"><b>{this.props.genreName}</b></div>
                    <div className="panel-body">{this.props.count}</div>
                    <div className="panel-footer">
                        {/*{this.props.joke.likes} Likes &nbsp;*/}
                        {/*<a onClick={this.like} className="btn btn-default">*/}
                        {/*    <span className="glyphicon glyphicon-thumbs-up" />*/}
                        {/*</a>*/}
                    </div>
                </div>
            </div>
        )
    }
}



ReactDOM.render(<App />, document.getElementById('app'));