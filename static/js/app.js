var app = angular.module('albumsApp', ['ngRoute', 'ngAnimate'])

.config(function($routeProvider) {
  $routeProvider
  .when('/', {
    controller: 'albumsCtrl',
    templateUrl: '/views/albums.html'
  })
  .when('/album/:name', {
    controller: 'albumCtrl',
    templateUrl: '/views/album.html'
  })
  .when('/image/:name', {
    controller: 'imageCtrl',
    templateUrl: '/views/image.html'
  })
});

app.controller('albumsCtrl', 
  function($scope, albumsFactory){
    $scope.albums = [];

    var handleSuccess = function(data, status) {
      $scope.albums = data.albums;
    };

    albumsFactory.getAlbums().success(handleSuccess);

    $scope.showCreateAlbum = false;

    $scope.$watch('searchText', function() {
      console.log("ctrl-searchText: "+$scope.searchText);
      $scope.$parent.searchText = $scope.searchText
    });

    $scope.edit = function(album) {
      event.preventDefault();
      console.log("edit: ")
      console.log(album)
      if($scope.albums == undefined) {
        return false
      }
      var aid = album.i
      $scope.albums[aid].ShowEdit = ! $scope.albums[aid].ShowEdit;
    };

    $scope.showEdit = function(album){
      if($scope.albums == undefined) {
        return false
      }
      var aid = album.i

      if($scope.albums[aid].ShowEdit == undefined) {
       $scope.albums[aid].ShowEdit = true;
      }
      return ! $scope.albums[aid].ShowEdit;
    };

    $scope.createShow = function() {
      $scope.showCreateAlbum = ! $scope.showCreateAlbum
    };

    var handleSuccessUpdate = function(data, status) {
      console.log(data)
    };
    var handleErrorUpdate = function(data, status) {
      console.log(data)
    };
    var handleSuccessCreate = function(data, status) {
      console.log(data)
    };
    var handleErrorCreate = function(data, status) {
      console.log(data)
    };

    // Create new album
    $scope.createSave = function(album) {
      event.preventDefault();
      albumsFactory.createAlbum(album).success(handleSuccessCreate).error(handleErrorCreate)
    };

    // Update album
    $scope.save = function(album) {
      albumsFactory.updateAlbum(album).success(handleSuccessUpdate).error(handleErrorUpdate)
    }
});

app.factory('albumsFactory', function($http){
  return {
    getAlbums: function() {
      return $http.get('/albums.json', { });
    },
    updateAlbum: function(album) {
      return $http.post('/album/update.json', { 
        data: album.album
      });
    },
    createAlbum: function(album) {
      return $http.post('/album/create.json', { 
        data: album.album
      });
    }
  };
});

app.controller('albumCtrl', 
  function($scope, albumFactory, albumsFactory, $routeParams) {
    $scope.images = [];
    $scope.selectedImages = [];
    $scope.albums = [];
    $scope.showAlbumMover = false;

    var handleSuccess = function(data, status) {
      $scope.title = data.Title;
      $scope.description = data.Description;
      $scope.images = data.images;
    };

    var name = $routeParams.name;
    albumFactory.getAlbum(name).success(handleSuccess);

    $scope.selImages = function() {
      return $scope.selectedImages.length > 0 ? true : false;
    }

    $scope.albumMove = function() {
      console.log("albumMove")
      console.log($scope.selectedImages)
      console.log($scope.albums)
    }

    $scope.albumMoveShow = function() {
      return $scope.showAlbumMover
    }

    $scope.albumMoveShower = function(shower) {
      event.preventDefault();
      if ($scope.showAlbumMover == true && shower.show == true){
        shower.show = false
      }
      
      $scope.showAlbumMover = shower.show
    }

    var handleSuccessAlbums = function(data, status) {
      $scope.albums = data.albums;
    }
    var handleSuccessMove = function(data, status) {
      console.log(data)
    }
    if ($scope.albums.length == 0) {
      albumsFactory.getAlbums().success(handleSuccessAlbums);
    }

    $scope.moveImages = function(album) {
      console.log(album)
      albumFactory.moveImages(album).success(handleSuccessMove);
    }
});

app.factory('albumFactory', function($http){
  return {
    getAlbum: function(name) {
      var url = '/album/'+name+'.json';
      return $http.get(url, { });
    },
    moveImages: function(album) {
      return $http.post('/album/move.json', { 
        data: album
      });
    },
  };
});

app.controller('imageCtrl', 
  function($scope, imageFactory, $routeParams) {
    var handleSuccess = function(data, status) {
      $scope.image = []
      $scope.image.src = data.Name
      $scope.image.fullsrc = data.Full
      $scope.image.tags = data.Tags
      $scope.loaded = false
    };

    var name = $routeParams.name;
    imageFactory.getImage(name).success(handleSuccess);
});

app.factory('imageFactory', function($http){
  return {
    getImage: function(name) {
      var url = '/image/'+name+'.json';
      return $http.get(url, { });
    }
  };
});

app.directive('createAlbumDialog', function() {
  return {
    restrict: 'E',
    templateUrl: '/views/albumcreate.html'
  };
});

app.directive('thumbViewer', function() {
  return {
    restrict: 'E',
    templateUrl: '/views/thumb.html'
  };
});

app.directive('imageonload', function($rootScope) {
  return {
    restrict: 'A',
    scope: true,
    link: function(scope, element) {
      element.bind('load', function() {
        scope.$parent.loaded = true
        scope.$apply();
      });
    }
  };
});

app.directive('multiSelect', function() {
  return {
    restrict: 'A',
    scope: {
      imageSelected: '&',
      selectedImages: '&'
    },
    controller: function($scope, $element, $attrs) {
      $scope.addSelected = function(id, name) {
        $scope.$parent.selectedImages.push({Name: name, Id: id});
        $scope.$apply();
        console.log($scope.$parent.selectedImages)
      };
      
      $scope.removeSelected = function(id, name) {
        var index = $scope.$parent.selectedImages.map(function(e) { return e.Name; }).indexOf(name);
        if (index > -1) {
          $scope.$parent.selectedImages.splice(index, 1);
          $scope.$apply();
          console.log($scope.$parent.selectedImages)
        }
      };
    },
    link: function(scope, element, attr, controllers ) {
      element.bind('click', function($event) {
        if ($event.ctrlKey == true) {
          event.preventDefault();
          var eid = element[0].attributes['data-id'].value
          var ename = element[0].attributes['data-value'].value
          if( typeof scope.imageSelected == "function" ) {
            scope.imageSelected = false;
          }
          scope.imageSelected = ! scope.imageSelected;
          element.toggleClass("selectedImage", scope.imageSelected)
          element[0].attributes['data-selected'].value = scope.imageSelected
          if (scope.imageSelected) {
            scope.addSelected(eid, ename)
          } else {
            scope.removeSelected(eid, ename)
          }
        }
      });   
    }
  };
});

app.directive('albumViewer', function($timeout) {
  return {
    restrict: 'E',
    transclude: true,
    replace: true,
    scope: {
      'album': '=',
      'albumIndex': '=',
      'searchText': '&',
    },
    link: function (scope, element) {
      scope.edit = scope.$parent.$parent.edit
      scope.showEdit = scope.$parent.$parent.showEdit
      scope.save = scope.$parent.$parent.save
      scope.createShow = scope.$parent.$parent.createShow
      scope.createSave = scope.$parent.$parent.createSave
    },
    templateUrl: '/views/albuminfo.html'
  };
});

app.directive('buttonAction', function() {
  return {
    restrict: 'A',
    scope: true,
    priority: 1001,
    controller: function($scope, $element, $attrs) {
    },
    link: function(scope, element, attr, controllers ) {
      element.bind('click', function($event) {
       scope.selectedImages = scope.$parent.selectedImages
       var oT = element.prop('offsetTop')
       var oL = element.prop('offsetLeft')﻿
       var myEl = angular.element( document.querySelector( '#albummover' ) );
       var noT = oT + 23 + "px"
       myEl.css('top', noT)
       var noL = oL - 95 + "px"
       myEl.css('left', noL)
       myEl.css('z-index', 99999)
      });   
    }
  };
});

var ModalCtrl = function ($scope, $modal, $log) {

  $scope.items = ['item_1', 'item_2', 'item_3'];
  console.log("modal")
  $scope.open = function () {
    console.log("modal open")
    var modalInstance = $modal.open({
      templateUrl: '/views/albumcreate.html',
      controller: ModalInstanceCtrl,
      resolve: {
        items: function () {
          return $scope.items;
        }
      }
    });

    modalInstance.result.then(function (selectedItem) {
      $scope.selected = selectedItem;
    }, function () {
      $log.info('Modal dismissed at: ' + new Date());
    });
  };
};

// Please note that $modalInstance represents a modal window (instance) dependency.
// It is not the same as the $modal service used above.

var ModalInstanceCtrl = function ($scope, $modalInstance, items) {

  $scope.items = items;
  $scope.selected = {
    item: $scope.items[0]
  };

  $scope.ok = function () {
    $modalInstance.close($scope.selected.item);
  };

  $scope.cancel = function () {
    $modalInstance.dismiss('cancel');
  };
};