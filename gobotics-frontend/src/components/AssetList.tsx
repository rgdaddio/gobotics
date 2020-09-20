import React from "react";

// TODO is this okay?
import Asset from "../ducks/assets/types";

import AssetCard from './AssetCard';

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