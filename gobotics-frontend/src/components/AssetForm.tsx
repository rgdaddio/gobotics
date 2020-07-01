import React from "react";

import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';

import { makeStyles } from '@material-ui/core/styles';
import Asset from "../types/Asset";
import { MouseEvent } from 'react';

const useStyles = makeStyles((theme) => ({
  root: {
    '& .MuiTextField-root': {
      margin: theme.spacing(1),
      width: '25ch',
    },
  },
}));

type AssetProps = {
    addAssetForm: Asset,
    handleChange: (event: MouseEvent<HTMLFormElement>) => void,
    onSubmit: (event: MouseEvent<HTMLButtonElement>) => void,
    showAddAssetForm: boolean
    assets: Array<Asset>
}

function ShowAddAssetForm(props: any) {
    const showForm = props.showAddAssetForm;
    console.log(props);
    if (showForm) {
       return ( <AssetForm 
        onSubmit={props.onSubmit}
        handleChange={props.handleChange}
        assets={props.assets} 
        addAssetForm={props.addAssetForm}
    /> );
    }
    return null;
  }

function AssetForm(props: any)  {
    const classes = useStyles();
    return(
        <div>
            <form className={classes.root} onChange={props.handleChange}>
                <div>
                    <TextField required id="assetName" name="name" label="Device Name" defaultValue="" autoComplete='off' value={props.addAssetForm.name} />   
                </div>   
                <div>
                    <TextField required id="assetLocation" name="ipaddr" label="IP Address" defaultValue="" autoComplete='off' value={props.addAssetForm.location} />
                    <Button  variant="contained" color="primary" onClick={props.onSubmit}>
                        Create Asset
                    </Button>
                </div>     
            </form>
        </div>
    );
}

export {
    AssetForm,
    ShowAddAssetForm 
}