import React from 'react'
import ReactDOM from 'react-dom'
import App from './common/component/app'
import './style.scss'
import { LoginProvider } from './common/component/login'
import { Sidebar } from './common/component/sidebar'
import * as serviceWorker from './serviceWorker'

ReactDOM.render(
  <React.StrictMode>
    <App>
      <LoginProvider>
        {(user) => {
          return <Sidebar user={user}/> 
        }}
      </LoginProvider>
    </App>
  </React.StrictMode>,
  document.body
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister()
