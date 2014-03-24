var app = angular.module('albumsApp', ['ngRoute'])

.config(function($routeProvider) {
  $routeProvider
  .when('/', {
    controller: 'albumController',
    templateUrl: '/views/albums.html'
  })
})

.controller('albumController', ['$scope', '$http',
  function ($scope, $http) {
    $http.get('/albums.json').success(function(data) {
      $scope.albums = data.albums;
    });

  $scope.showCreateAlbum = false

  $scope.edit = function(album) {
    event.preventDefault();
    if($scope.albums == undefined) {
      return false
    }
    var aid = album.i
    $scope.albums[aid].ShowEdit = ! $scope.albums[aid].ShowEdit;
  }
  $scope.showEdit = function(album){
    event.preventDefault();
    if($scope.albums == undefined) {
      return false
    }
    var aid = album.i

    if($scope.albums[aid].ShowEdit == undefined) {
     $scope.albums[aid].ShowEdit = true;
    }
    return ! $scope.albums[aid].ShowEdit;
  }
  $scope.save = function(album) {
    $http({
      url: '/album/update.json',
      method: "POST",
      data: album.album
    })
    .success(function (data, status, headers, config) {
      console.log("success")
      $scope.edit(album)
    }).error(function (data, status, headers, config) {
      console.log("error")
    });
  }
  $scope.createShow = function() {
    $scope.showCreateAlbum = ! $scope.showCreateAlbum
  }
  $scope.createSave = function(album) {
    $http({
      url: '/album/create.json',
      method: "POST",
      data: album
    })
    .success(function (data, status, headers, config) {
      console.log("success")
      $scope.createShow()
      $location.url('/albums#/')
    }).error(function (data, status, headers, config) {
      console.log("error")
    });
  }
}]);

app.directive('albumViewer', function() {
  return {
    restrict: 'E',
    scope: {
      'album': '=',
      'albumIndex': '='
    },
    link: function (scope, elem, attrs) {
    },
    templateUrl: '/views/album.html'
  };
});

app.directive('createAlbumDialog', function() {
  return {
    restrict: 'E',
    templateUrl: '/views/albumcreate.html'
  };
});
