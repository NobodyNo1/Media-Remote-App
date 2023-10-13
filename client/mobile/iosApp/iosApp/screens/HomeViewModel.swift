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
    let isLoading   : Bool      = false
    let fail        : String?   = nil
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
        remoteMediaController.playPause { (err : Error?) -> Void in
            
        }
    }
    
    func updateVolume(increase: Bool) {
        //async await call
        remoteMediaController.updateVolume(shouldIncrease: increase) { (err : Error?) -> Void in
            
        }
    }
}
