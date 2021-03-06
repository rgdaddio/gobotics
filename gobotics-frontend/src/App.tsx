import React, {Component} from "react";
import AssetList from "./components/AssetList";
import {ShowAddAssetForm, AssetForm} from "./components/AssetForm";
import Asset from "./types/Asset";
import Button from '@material-ui/core/Button';
import { MouseEvent } from 'react';

// todo super(props) ?


// hey quick idea for your "showing data" thing that you might find useful
//as a POF you could have a component that has both the vanilla card + info card stacked

//TODO function components with state how in react
// React State and props together 

//TODO remove all props:any
// https://www.typescriptlang.org/docs/handbook/react.html
// export interface Props {
// https://stackoverflow.com/questions/52735288/why-does-parameter-props-implicitly-has-an-any-type
//https://stackoverflow.com/questions/47561848/property-value-does-not-exist-on-type-readonly



type AppProps = {}

type AppState = { 
  showAddAssetForm: boolean,
  addAssetForm: Asset,
  assets: Array<Asset>
};

class App extends Component<AppProps, AppState>{
    constructor(props:AppProps) {
        super(props);
        fetch('/client/devices').
        then(res => res.json()).
        then(list => {
            const entries = list.map( (x: any) => ({
                'id': x.name,
                'name': x.name,
                'mac_address': x.mac_address,
                'ip_address': x.ip_address
            }));
            this.setState({assets: this.state.assets.concat(entries)})
        })
        this.state = {
            'showAddAssetForm': false,
            'addAssetForm': {
                'id' : '',
                'name': '',
                'ip_address': '',
                'mac_address': ''
            },
            assets: []
        }
    }

    handleAssetFormChange = (event : React.ChangeEvent<HTMLFormElement>) => {
        const addAssetFormState = this.state.addAssetForm;
        const { name, value } = event.target;
        addAssetFormState[name] = value;
        this.setState({ addAssetForm: addAssetFormState });
      };

    onSubmit = (event: MouseEvent<HTMLButtonElement>) => {
        event.preventDefault();
        if (this.state.addAssetForm.name == "" && this.state.addAssetForm.ip_address == "")
            return;
        else {
            const newAsset: Asset = {
              "id": String(Math.random()),
              "name": this.state.addAssetForm.name,
              "mac_address": this.state.addAssetForm.mac_address,
              "ip_address": this.state.addAssetForm.ip_address
            };
            this.setState({
                assets: this.state.assets.concat(newAsset),
                // TODO ask rohan
                addAssetForm: {name: "", id: "", ip_address: "", mac_address: ""}
            });
        }
    }

    // bind click to some state variable 
    // conditonal check to show asset form
    render() {
        return (
            <div className="assetContainer">
                <Button onClick={event => this.setState({showAddAssetForm: !this.state.showAddAssetForm})} variant="contained" color="primary">
                    New Asset
                </Button>

                <ShowAddAssetForm
                    showAddAssetForm={this.state.showAddAssetForm}
                    onSubmit={this.onSubmit}
                    handleChange={this.handleAssetFormChange}
                    assets={this.state.assets} 
                    addAssetForm={this.state.addAssetForm}/>

                <AssetList 
                    onSubmit={this.onSubmit}
                    assets={this.state.assets} 
                    />
            </div>
            
        );
    }
}

export default App;