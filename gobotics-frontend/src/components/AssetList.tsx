import React from "react";
import Asset from "../types/Asset";

import {Card, CardContent} from '@material-ui/core';

function AssetCard(props: any) {
    return (
        <Card variant="outlined">
        <CardContent>
            <h4>{props.asset.name}</h4>
            <div>id: {props.asset.id}</div>
            <div>Status: {props.asset.device_status} </div>
            <div>Location: {props.asset.location} </div>
        </CardContent>
    </Card>
        );
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