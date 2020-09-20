# gobotics-frontend
A proof of concept frontend / portal for gobotics to help learn about React/Redux.
`gobotics-frontend` usees the React framework and Redux for statemanagment. 

# Building
A makefile is provided, so you don't have to remember the npm commands. To run the js code as is, use `make start`. This will
just start the npm server on `localhost:3000`. This is the most convenient for local development and debugging. 

For caching and performance reasons, the `gobotics-server` uses a production build of the UI. Run `make create` to run the production build. This needs to be run everytime you make a change to the frontend. Once finished `gobotics-server` knows where to look for the static files and will handle the caching. 

# Reference links:
* https://react-redux.js.org/introduction/quick-start
* https://redux.js.org/basics/usage-with-react
React with typescript:
* https://www.typescriptlang.org/docs/handbook/react.html
For redux, the frontend uses the `ducks` for of file organization. 
Articles used for reference are: 
* https://www.freecodecamp.org/news/scaling-your-redux-app-with-ducks-6115955638be/
* https://github.com/FortechRomania/react-redux-complete-example
