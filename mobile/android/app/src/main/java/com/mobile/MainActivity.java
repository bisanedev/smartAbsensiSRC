package com.mobile;
import android.os.Bundle;  // splashscreen
import com.facebook.react.ReactActivity;
import org.devio.rn.splashscreen.SplashScreen; // splashscreen

public class MainActivity extends ReactActivity {

  /**
   * Returns the name of the main component registered from JavaScript. This is used to schedule
   * rendering of the component.
   */
  @Override
  protected void onCreate(Bundle savedInstanceState) {    
    SplashScreen.show(this, R.style.SplashScreenTheme);
    super.onCreate(savedInstanceState);
  }  
  protected String getMainComponentName() {
    return "mobile";
  }
}
