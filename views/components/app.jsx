import React from "react"
import ReactDOM from "react-dom"
import styled from "styled-components"
import {postJson, colorScale, appdata} from "../constants/AppConstants"
import {BrowserRouter as Router, Link, Route} from "react-router-dom"
import WorldMap from "./viz/WorldMap"
import { range } from 'd3-array'


// class App extends React.Component {
//
//     render() {
//         this.loggedIn = true;
//         if (this.loggedIn) {
//             return (<LoggedIn />);
//         } else {
//             return (<Home />);
//         }
//     }
// }

// class Home extends React.Component {
//     render() {
//         return (
//             <div className="container">
//                 <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
//                     <h1>Spotify Insights</h1>
//                     <p>Insight into your personal Spotify Account</p>
//                     <p>Sign in to get access </p>
//                     /*<a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>*/
//                 </div>
//             </div>
//         )
//     }
// }

appdata
    .forEach((d,i) => {
        const offset = Math.random()
        d.launchday = i
        d.data = range(30).map((p,q) => q < i ? 0 : Math.random() * 2 + offset)
    })

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

                                <Link to={`/graph`}>
                                    <button className="btn btn-primary">Graph</button>
                                </Link>
                            </div>
                        </div>
                    </Row>

                    <Route path="/genre-count" component={GenreCountContainer} />
                    <Route path="/graph" component={Graph} />

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
                    {  Object.keys(this.state.genreCounts).map((key, index) => (
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

class Graph extends React.Component {
    constructor(props) {
        super(props)
        this.onResize = this.onResize.bind(this)
        this.onHover = this.onHover.bind(this)
        this.onBrush = this.onBrush.bind(this)
        this.state = { 
            screenWidth: 1000,
            screenHeight: 500,
            hover: "none", 
            brushExtent: [0,40] 
        }

    }

    onResize() {
        this.setState({ 
            screenWidth: window.innerWidth + 1000,
            screenHeight: window.innerHeight - 120 + 1000})
    }

    onHover(d) {
        this.setState({ hover: d.id })
    }

    onBrush(d) {
        this.setState({ brushExtent: d })
    }

    componentDidMount() {
        window.addEventListener('resize', this.onResize, false)
        this.onResize()
    }

    render() {
        const filteredAppdata = appdata
            .filter((d,i) => d.launchday >= this.state.brushExtent[0] && d.launchday <= this.state.brushExtent[1])
        return (
            <div className="row">
                <WorldMap hoverElement={this.state.hover} onHover={this.onHover} colorScale={colorScale} data={filteredAppdata} size={[this.state.screenWidth / 2, this.state.screenHeight / 2]} />
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