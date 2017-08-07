'use strict';

angular.module('app.main', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/main', {
    templateUrl: 'views/main/main.html',
    controller: 'MainCtrl'
  });
}])

.controller('MainCtrl', function($scope, $http) {
    $scope.addresses = [];
    $http.get(urlapi + 'lasttx')
        .then(function(data, status, headers, config) {
            console.log('data success');
            console.log(data);

            $scope.addresses = data.data;
        }, function(data, status, headers, config) {
            console.log('data error');
        });
});
