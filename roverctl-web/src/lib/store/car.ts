// /**
//  * Car state that updates so quickly that it should be its own store
//  */

// import { writable } from "svelte/store";

// type Vector = {
//   x: number;
//   y: number;
//   z: number;
// };

// type CarState = {
//   gyroscope: Vector | null;
//   battery: {
//     voltage: number;
//     warnVoltage: number;
//     killVoltage: number;
//     timestamp: number; // so that we know when the last update was
//   } | null;
//   lapTimes: number[]; // durations of laptimes
// };

// const createCarStore = () => {
//   const store = writable<CarState>({
//     gyroscope: null,
//     battery: null,
//     lapTimes: [],
//   });
//   const { subscribe, update, set } = store;

//   return {
//     // Required functions
//     subscribe,
//     update,
//     set,
//   };
// };

// type CarStore = ReturnType<typeof createCarStore>;

// const carStore = createCarStore();

// export { carStore };
// export type { CarStore, CarState };
