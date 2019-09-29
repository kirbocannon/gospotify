import React from "react";
import {appdata, colorScale} from "./constants/AppConstants";
import {Link} from "react-router-dom";
import WorldMap from "./viz/WorldMap"
import { range } from 'd3-array'

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

    componentWillUnmount() {
        window.removeEventListener('resize', this.onResize, false)

    }

    // componentWillUnmount() {
    //     let el = ReactDom.findDOMNode(this);
    //     d3Chart.destroy(el);
    // }

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

appdata
    .forEach((d,i) => {
        const offset = Math.random()
        d.launchday = i
        d.data = range(30).map((p,q) => q < i ? 0 : Math.random() * 2 + offset)
    })

<Link to={`/graph`}>
    <button className="btn btn-primary">Graph</button>
</Link>

<Route path="/graph" component={Graph} />



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
                    <h1>Spotify Insights</h1>
                    <p>Insight into your personal Spotify Account</p>
                    <p>Sign in to get access </p>
                    /*<a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>*/
                </div>
            </div>
        )
    }
}