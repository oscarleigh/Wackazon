//
//  OnboardingView.swift
//  Wackazon
//
//  Created by Oscar Leigh on 02/04/2026.
//

import SwiftUI

struct OnboardingView: View {
    @State private var viewModel = OnBoardingViewModel()
    
    var body: some View {
        NavigationStack {
            VStack (spacing: 20) {
                Text("Wackazon")
                    .font(Font.largeTitle.bold())
                
                Text("A Practice project by Oscar & Matt")
                    .opacity(0.75)
                
                Button() {
                    viewModel.goToSignIn()
                } label: {
                    Text("Sign in")
                        .frame(maxWidth: .infinity)
                }
                .buttonStyle(.glassProminent)
                .padding(.horizontal, 50)
                
                
                Button() {
                    viewModel.goToRegister()
                } label: {
                    Text("Register")
                        .frame(maxWidth: .infinity)
                }
                .buttonStyle(.glass)
                .padding(.horizontal, 50)
                
                Button() {
                    viewModel.goToHomePage()
                } label: {
                    Text("Contine as Guest")
                }
            }
            .navigationDestination(isPresented: $viewModel.navigateToSignIn) {
                SignInView()
            }
            .navigationDestination(isPresented: $viewModel.navigateToRegister) {
                RegisterView()
            }
        }
    }
}

#Preview {
    OnboardingView()
}
