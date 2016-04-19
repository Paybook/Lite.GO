var gulp = require('gulp');
var flatten = require('gulp-flatten');
var bower = require('bower-files')();
var angularFilesort = require('gulp-angular-filesort');
var inject = require('gulp-inject');
var series = require('stream-series');
var uglify = require('gulp-uglify');
var concat = require('gulp-concat');

gulp.task('default', function() {

    gulp.src(bower.ext('js').files)
    .pipe(concat('lib.min.js'))
    .pipe(uglify())
    .pipe(gulp.dest('static/js/dependencies'));

    gulp.src('bower_components/**/*.min.css')
    .pipe(flatten())
    .pipe(gulp.dest('static/css'));


    var dependencies = gulp.src(['./static/js/dependencies/lib.min.js','./static/**/*css'], {read: false});
    var app = gulp.src(['./static/js/app.js'], {read: false});
    var controllers = gulp.src(['./static/js/controllers/*.js'], {read: false});

    gulp.src('./views/inc/layout.tpl')
    .pipe(inject(series(dependencies, app, controllers))) // This will always inject vendor files before app files
    .pipe(gulp.dest('./views/inc'));

});