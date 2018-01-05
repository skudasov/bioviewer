import {RECEIVE_PROFILE, RECEIVE_PROFILE_ERR} from '../constants/ActionTypes'

export const profile_reducer = (state = null, action) => {
    switch (action.type) {
        case RECEIVE_PROFILE:
            return action.singleProfile;
        case RECEIVE_PROFILE_ERR:
            return action.err;
        default:
            return state
    }
};