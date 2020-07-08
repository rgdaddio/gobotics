import Asset from "./types";

// Actions
const ADD_ASSET = 'ADD_ASSET'
const ADD_ASSETS = 'ADD_ASSETS'

// Action Creator typescript interfaces
export interface AddAssetAction {
    type: typeof ADD_ASSET,
    payload: Asset
}

export interface AddAssetsAction {
    type: typeof ADD_ASSETS,
    payload: Asset[]
}

// Action Creators
const addAsset = (asset:Asset): AddAssetAction => ({
      type: ADD_ASSET,
      payload: asset
      //TODO generate a unique id every time something new is created?
      //https://redux.js.org/basics/actions
});

const addAssets =(assets:Asset[]): AddAssetsAction => ({
      type: ADD_ASSETS,
      payload: assets
      //TODO generate a unique id every time something new is created?
      //https://redux.js.org/basics/actions
});
 
 export default {
    ADD_ASSET,
    ADD_ASSETS,
    addAsset,
    addAssets
 }