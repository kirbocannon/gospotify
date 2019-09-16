import { scaleThreshold } from 'd3-scale'
import {geoCentroid} from "d3-geo"
import worlddata from "../components/viz/world"

const postJson = {
    method: 'POST',
    headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
    }
};

const colorScale = scaleThreshold().domain([5,10,20,30]).range(["#75739F", "#5EAFC6", "#41A368", "#93C464"]);

const appdata = worlddata.features
    .filter(d => geoCentroid(d)[0] < -20);

export {
    postJson,
    colorScale,
    appdata
}