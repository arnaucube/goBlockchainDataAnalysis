'use strict';

angular.module('app.beta', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/beta', {
            templateUrl: 'views/beta/beta.html',
            controller: 'BetaCtrl'
        });
    }])

    .controller('BetaCtrl', function($scope, $http, $routeParams) {
        
    });
