import { connect } from 'react-redux'
import Asset from "../ducks/assets/types";
import AssetList from '../components/AssetList'

const getAssets = (assets: Asset[], filter = "ALL") => {
    switch(filter) {
        // can filter down the list by type here maybe?
        default:
            return assets;
    }
}
const mapStateToProps = state => {
    return {
      todos: getAssets(state.assets)
    }
  }

const mapDispatchToProps = dispatch => ({
    //TODO
})
  
export default connect(mapStateToProps, mapDispatchToProps)(AssetList)
