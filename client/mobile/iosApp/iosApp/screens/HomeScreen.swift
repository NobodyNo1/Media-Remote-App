//
//  HomeScreen.swift
//  iosApp
//
//  Created by Bekarys on 14.10.2023.
//  Copyright © 2023 orgName. All rights reserved.
//

import Foundation
import SwiftUI

struct HomeScreen: View {

    @ObservedObject var homeViewModel = HomeViewModel()

    var body: some View {
        VStack {
            ProgressView().opacity(
                homeViewModel.uiState.isLoading ? 1.0 : 0.0
            )
            Spacer().frame(height: 26)
            
            VolumeControlView(updateVolume: { (increase: Bool) -> Void in
                homeViewModel.updateVolume(increase: increase)
            })
            Spacer().frame(height: 26)
            Text("Media Keys")
            Spacer().frame(height: 8)
            
            Button("Play/Pause", action: {
                homeViewModel.playPause()
            })
            Spacer().frame(height: 26)


            if(homeViewModel.uiState.fail != nil) {
                Text(homeViewModel.uiState.fail!)
            }
        
        }
    }
    
    struct VolumeControlView: View {
        
        let updateVolume: (Bool)->Void
        
        init(updateVolume: @escaping (Bool)->Void) {
            self.updateVolume = updateVolume
        }
        
        var body: some View {
            Text("Volume Control")
            Spacer().frame(height: 8)
            
            HStack {
                Spacer()
                Button("Decrease", action: {
                    updateVolume(false)
                    //homeViewModel.updateVolume(increase: false)
                })
                Spacer()
                Button("Increase", action: {
                    updateVolume(true)
                    //homeViewModel.updateVolume(increase: true)
                })
                Spacer()
            }
            Spacer().frame(height: 8)
        }
    }
}
