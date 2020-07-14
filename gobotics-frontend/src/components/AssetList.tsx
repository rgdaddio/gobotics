import React from "react";
import Asset from "../ducks/assets/types";
import {Card, CardContent} from '@material-ui/core';

// Presentational Components

function AssetCard(props: any) {
    return      (
        <Card variant="outlined">
        <CardContent>
            <h4>{props.asset.name}</h4>
            <div>Id: {props.asset.id}</div>
            <div>IP Address: {props.asset.ip_address} </div>
            <div>MAC Address: {props.asset.mac_address} </div>
        </CardContent>
    </Card>);
}

function AssetList(props: any) {
    return (
        <React.Fragment>
            <div>
            <h1>Sensors:</h1>
            {props.assets.map((asset: Asset) => (
                <div  key={asset.id} > 
                    <AssetCard asset={asset} />
                </div>
            ))}
            </div>
        </React.Fragment>
    );
}



export default AssetList; 