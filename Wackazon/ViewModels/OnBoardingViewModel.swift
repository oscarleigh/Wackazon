//
//  OnBoardingViewModel.swift
//  Wackazon
//
//  Created by Oscar Leigh on 02/04/2026.
//

import SwiftUI

@Observable
class OnBoardingViewModel {
    var navigateToSignIn = false
    var navigateToRegister = false
    var navigateToHomePage = false
    
    func goToSignIn() {
        navigateToSignIn = true
    }
    
    func goToRegister() {
        navigateToRegister = true
    }
    
    func goToHomePage() {
        navigateToHomePage = true
    }
}
