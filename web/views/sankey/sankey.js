'use strict';

angular.module('app.sankey', ['ngRoute'])

    .config(['$routeProvider', function($routeProvider) {
        $routeProvider.when('/sankey', {
            templateUrl: 'views/sankey/sankey.html',
            controller: 'SankeyCtrl'
        });
    }])

    .controller('SankeyCtrl', function($scope, $http, $routeParams) {
        

    });
