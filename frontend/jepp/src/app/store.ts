import { configureStore } from "@reduxjs/toolkit";

import configReducer from "../features/config/configSlice"

export const store = configureStore({
	reducer: {
		config: configReducer,
	},
});

// Infer the `RootState` and `AppDispatch` types from the store itself
export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
