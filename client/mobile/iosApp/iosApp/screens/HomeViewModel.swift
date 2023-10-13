//
//  HomeViewModel.swift
//  iosApp
//
//  Created by Bekarys on 14.10.2023.
//  Copyright © 2023 orgName. All rights reserved.
//

import Foundation
import shared


struct HomeUiState {
    let isLoading   : Bool
    let fail        : String?
    
    
    init() {
        self.isLoading  = false
        self.fail       = nil
    }
    
    init(
        isLoading: Bool,
        fail     : String?  = nil
    ) {
        self.isLoading  = isLoading
        self.fail       = fail
    }

    func changeValues(
        isLoading: Bool?    = nil,
        fail     : String?  = nil
    ) -> HomeUiState {
        return HomeUiState(isLoading: isLoading ?? self.isLoading, fail: fail ?? self.fail)
    }
}

final class HomeViewModel: ObservableObject {
    @Published var uiState = HomeUiState()
    let remoteMediaController = RemoteMediaController()
//    @Published var weather = WeatherResponse.empty()
//    @Published var city: String = "Bogotá" {
//        didSet {
//            getLocation()
//        }
//    }
//
//    private lazy var dateFormatter: DateFormatter = {
//        let formatter = DateFormatter()
//        formatter.dateStyle = .full
//        return formatter
//    }()
//
//    private lazy var dayFormatter: DateFormatter = {
//        let formatter = DateFormatter()
//        formatter.dateFormat = "EEE"
//        return formatter
//    }()
//
//    private lazy var timeFormatter: DateFormatter = {
//        let formatter = DateFormatter()
//        formatter.dateFormat = "hh a"
//        return formatter
//    }()
    
    init() {
//        getLocation()
    }
    
    
    func playPause() {
        //async await call
        uiState = uiState.changeValues(isLoading: true)
        remoteMediaController.playPause { (err : Error?) -> Void in

            DispatchQueue.main.async {
                self.uiState =  self.uiState.changeValues(isLoading: false)
            }
        }
    }
    
    func updateVolume(increase: Bool) {
        //async await call
        uiState = uiState.changeValues(isLoading: true)
        remoteMediaController.updateVolume(shouldIncrease: increase) { (err : Error?) -> Void in

            DispatchQueue.main.async {
                self.uiState =  self.uiState.changeValues(isLoading: false)
            }
        }
    }
}
