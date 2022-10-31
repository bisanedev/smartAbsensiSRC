const mix = require('laravel-mix');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */
// admin
mix.js('resources/js/apps.js', 'public/assets/js');
mix.sass('resources/sass/apps.scss', 'public/assets/css');
// copy asset
mix.copyDirectory('resources/assets', 'public/assets');