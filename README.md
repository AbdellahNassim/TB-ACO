# Trust-Based Ant Colony Optimization (TB-ACO) 🐜✨

### **Little Kids, Are You Still Sleeping? :)**
Before you proceed, take a moment to say **Alhamdulillah** and reflect on the importance of wisdom, trust, and learning from mistakes. This is not just an algorithm—it’s a **mentality shift**. If you don’t believe and seek knowledge sincerely, the code won’t reveal its secrets to you. 😏

---

## 🌍 **Introduction**
TB-ACO is a **metaheuristic algorithm** that combines **Ant Colony Optimization (ACO) with Reinforcement Learning**, inspired by the principles of **trust and self-reliance in Allah (SWT)**. Unlike traditional ACO, where ants follow pheromone trails, TB-ACO introduces an agent (KA: Knowledgeable Ant) that **learns from mistakes, does not trust others blindly, and eventually leads the swarm to a global optimum**. 

### **🔑 Key Features**
✅ **Trust-Based Learning** – KA follows its own path, avoiding mistakes it has already learned from.  
✅ **Guiding Others** – Other ants observe KA’s success and follow its lead over time.  
✅ **Never-Declining Trust** – Trust grows exponentially as KA continues to make optimal choices. Even mistakes are learning opportunities. 
✅ **Adaptive Trust Floor** – Mistakes do **not** decrease trust; rather, they contribute to an ever-increasing minimum trust threshold. This ensures KA’s influence never diminishes over time.  

---

## 📜 The Math Behind TB-ACO

### **1. Trust Evolution**
Trust in KA is modeled as a function of its past actions. 

#### **Trust Function** (No Decay Version)
Let:
- \( T(t) \) be the trust at time \( t \)
- \( T_0 \) be the initial trust level
- \( \alpha \) be the learning rate of trust
- \( R(t) \) be KA’s cumulative success count at \( t \)

We define:
\[ T(t) = T_0 + (1 - T_0) (1 - e^{-\alpha R(t)}) \]
where **trust in KA only grows** as more successes are observed. Trust does not decay because every mistake is a learning opportunity.

---

## 🤖 **Algorithm Overview**
```rust
use rand::Rng;

struct Colony {
    trust: f64,
    trust_min: f64,
    success_rate: f64,
}

impl Colony {
    fn update_trust(&mut self, success: bool) {
        let alpha = 0.1; // Learning rate
        let gamma = 0.02; // Adaptive threshold growth factor
        
        if success {
            self.success_rate += 1.0;
            self.trust += (1.0 - self.trust) * alpha;
        }
        self.trust_min = (0.1 + 0.02 * self.success_rate).min(1.0);
        self.trust = self.trust.max(self.trust_min);
    }
}

fn main() {
    let mut rng = rand::thread_rng();
    let mut colony = Colony { trust_min: 0.1, trust: 0.1, success_rate: 0.0 };
    
    println!("Little kids, are you still sleeping? Say 'Alhamdulillah' first before executing this! 😏");
    
    for t in 1..=100 {
        let success = rng.gen::<f64>() < 0.7; // 70% success rate for KA
        colony.update_trust(success);
        println!("Iteration {}: Trust = {:.4}, TrustMin = {:.4}", t, colony.trust, colony.trust_min);
    }
}
