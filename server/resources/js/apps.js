window._ = require('lodash');

/**
 * We'll load the axios HTTP library which allows us to easily issue requests
 * to our Laravel back-end. This library automatically handles sending the
 * CSRF token as a header based on the value of the "XSRF" token cookie.
 */
try {
    window.Popper = require('popper.js').default;   
    require('bootstrap');
	window.$ = window.jQuery = require('jquery');				
} catch (e) {}

window.jwt_decode = require('jwt-decode');
window.axios = require('axios');
//vuejs
var token = window.localStorage.getItem('token');

if (token) {
	window.axios.defaults.headers.common['Authorization'] = token;    
}

require('./vue');