import React from "react";

import {Card, CardContent} from '@material-ui/core';

// TODO is this okay?
import Asset from "../ducks/assets/types";
import PropTypes from 'prop-types'


// ???
const onClick = () => ({ console.log(" asset was clicked")})
type AssetProps = {
    asset: Asset,
    onClick : onClick
}

function AssetCard(props: AssetProps) {
    return      (
        <Card variant="outlined" onClick={props.onClick}>
            <CardContent>
                <h4>{props.asset.name}</h4>
                <div>Id: {props.asset.id}</div>
                <div>IP Address: {props.asset.ip_address} </div>
                <div>MAC Address: {props.asset.mac_address} </div>
            </CardContent>
        </Card>
    );
}

// ????
AssetCard.propTypes = {
    onClick: PropTypes.func.isRequired,
    completed: PropTypes.bool.isRequired,
    text: PropTypes.string.isRequired
  }

export default AssetCard