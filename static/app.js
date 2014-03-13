var app = angular.module('albumsApp', ['ngRoute'])

.config(function($routeProvider) {
	$routeProvider.when('/', {
		controller: "albumController",
		templateUrl: "/views/album.html"
	})
})

.controller('albumController', ['$scope', '$http',
  function ($scope, $http) {
    $http.get('/albums.json').success(function(data) {
      $scope.albums = data.albums;
    });
}]);